package db

import (
	"context"
	"fmt"
	"time"

	"code.cloudfoundry.org/lager/lagerctx"
	sq "github.com/Masterminds/squirrel"
	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/creds"
	"github.com/concourse/concourse/atc/db/lock"
)

//go:generate counterfeiter . Checkable

type Checkable interface {
	PipelineRef

	Name() string
	TeamID() int
	ResourceConfigScopeID() int
	TeamName() string
	Type() string
	Source() atc.Source
	Tags() atc.Tags
	CheckEvery() string
	CheckTimeout() string
	LastCheckEndTime() time.Time
	CurrentPinnedVersion() atc.Version

	HasWebhook() bool

	SetResourceConfig(
		atc.Source,
		atc.VersionedResourceTypes,
	) (ResourceConfigScope, error)

	CheckPlan(atc.Version, time.Duration, time.Duration, ResourceTypes) atc.CheckPlan
	CreateBuild(context.Context, bool) (Build, bool, error)

	SetCheckSetupError(error) error
}

//go:generate counterfeiter . CheckFactory

type CheckFactory interface {
	TryCreateCheck(context.Context, Checkable, ResourceTypes, atc.Version, bool) (Build, bool, error)
	Resources() ([]Resource, error)
	ResourceTypes() ([]ResourceType, error)
}

type checkFactory struct {
	conn        Conn
	lockFactory lock.LockFactory

	secrets       creds.Secrets
	varSourcePool creds.VarSourcePool

	planFactory atc.PlanFactory

	defaultCheckTimeout             time.Duration
	defaultCheckInterval            time.Duration
	defaultWithWebhookCheckInterval time.Duration
}

func NewCheckFactory(
	conn Conn,
	lockFactory lock.LockFactory,
	secrets creds.Secrets,
	varSourcePool creds.VarSourcePool,
	defaultCheckTimeout time.Duration,
	defaultCheckInterval time.Duration,
	defaultWithWebhookCheckInterval time.Duration,
) CheckFactory {
	return &checkFactory{
		conn:        conn,
		lockFactory: lockFactory,

		secrets:       secrets,
		varSourcePool: varSourcePool,

		planFactory: atc.NewPlanFactory(time.Now().Unix()),

		defaultCheckTimeout:             defaultCheckTimeout,
		defaultCheckInterval:            defaultCheckInterval,
		defaultWithWebhookCheckInterval: defaultWithWebhookCheckInterval,
	}
}

func (c *checkFactory) TryCreateCheck(ctx context.Context, checkable Checkable, resourceTypes ResourceTypes, from atc.Version, manuallyTriggered bool) (Build, bool, error) {
	logger := lagerctx.FromContext(ctx)

	var err error

	parentType, found := resourceTypes.Parent(checkable)
	if found {
		if parentType.Version() == nil {
			return nil, false, fmt.Errorf("resource type '%s' has no version", parentType.Name())
		}
	}

	interval := c.defaultCheckInterval
	if checkable.HasWebhook() {
		interval = c.defaultWithWebhookCheckInterval
	}
	if every := checkable.CheckEvery(); every != "" {
		interval, err = time.ParseDuration(every)
		if err != nil {
			return nil, false, fmt.Errorf("check interval: %s", err)
		}
	}

	if !manuallyTriggered && time.Now().Before(checkable.LastCheckEndTime().Add(interval)) {
		// skip creating the check if its interval hasn't elapsed yet
		return nil, false, nil
	}

	timeout := c.defaultCheckTimeout
	if to := checkable.CheckTimeout(); to != "" {
		timeout, err = time.ParseDuration(to)
		if err != nil {
			return nil, false, fmt.Errorf("check timeout: %s", err)
		}
	}

	checkPlan := checkable.CheckPlan(from, interval, timeout, resourceTypes.Filter(checkable))

	plan := c.planFactory.NewPlan(checkPlan)

	build, created, err := checkable.CreateBuild(ctx, manuallyTriggered)
	if err != nil {
		return nil, false, fmt.Errorf("create build: %w", err)
	}

	if !created {
		return nil, false, nil
	}

	_, err = build.Start(plan)
	if err != nil {
		return nil, false, fmt.Errorf("start build: %w", err)
	}

	logger.Info("created-build", build.LagerData())

	_, err = build.Reload()
	if err != nil {
		return nil, false, fmt.Errorf("reload build: %w", err)
	}

	return build, true, nil
}

func (c *checkFactory) Resources() ([]Resource, error) {
	var resources []Resource

	rows, err := resourcesQuery.
		Where(sq.Eq{"p.paused": false}).
		RunWith(c.conn).
		Query()

	if err != nil {
		return nil, err
	}

	defer Close(rows)

	for rows.Next() {
		r := newEmptyResource(c.conn, c.lockFactory)
		err = scanResource(r, rows)
		if err != nil {
			return nil, err
		}

		resources = append(resources, r)
	}

	return resources, nil
}

func (c *checkFactory) ResourceTypes() ([]ResourceType, error) {
	var resourceTypes []ResourceType

	rows, err := resourceTypesQuery.
		RunWith(c.conn).
		Query()

	if err != nil {
		return nil, err
	}

	defer Close(rows)

	for rows.Next() {
		r := newEmptyResourceType(c.conn, c.lockFactory)
		err = scanResourceType(r, rows)
		if err != nil {
			return nil, err
		}

		resourceTypes = append(resourceTypes, r)
	}

	return resourceTypes, nil
}
