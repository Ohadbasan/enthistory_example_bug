// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"entdemo/ent/migrate"

	"entdemo/ent/testentity"
	"entdemo/ent/testentityhistory"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/flume/enthistory"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// TestEntity is the client for interacting with the TestEntity builders.
	TestEntity *TestEntityClient
	// TestEntityHistory is the client for interacting with the TestEntityHistory builders.
	TestEntityHistory *TestEntityHistoryClient
	// historyActivated determines if the history hooks have already been activated
	historyActivated bool
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.TestEntity = NewTestEntityClient(c.config)
	c.TestEntityHistory = NewTestEntityHistoryClient(c.config)
}

// withHistory adds the history hooks to the appropriate schemas - generated by enthistory
func (c *Client) WithHistory() {
	if !c.historyActivated {

		// TestEntity hooks
		c.TestEntity.Use(enthistory.HistoryTriggerInsert[*TestEntityMutation]())

		c.historyActivated = true
	}
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		TestEntity:        NewTestEntityClient(cfg),
		TestEntityHistory: NewTestEntityHistoryClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:               ctx,
		config:            cfg,
		TestEntity:        NewTestEntityClient(cfg),
		TestEntityHistory: NewTestEntityHistoryClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		TestEntity.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.TestEntity.Use(hooks...)
	c.TestEntityHistory.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.TestEntity.Intercept(interceptors...)
	c.TestEntityHistory.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *TestEntityMutation:
		return c.TestEntity.mutate(ctx, m)
	case *TestEntityHistoryMutation:
		return c.TestEntityHistory.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// TestEntityClient is a client for the TestEntity schema.
type TestEntityClient struct {
	config
}

// NewTestEntityClient returns a client for the TestEntity from the given config.
func NewTestEntityClient(c config) *TestEntityClient {
	return &TestEntityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `testentity.Hooks(f(g(h())))`.
func (c *TestEntityClient) Use(hooks ...Hook) {
	c.hooks.TestEntity = append(c.hooks.TestEntity, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `testentity.Intercept(f(g(h())))`.
func (c *TestEntityClient) Intercept(interceptors ...Interceptor) {
	c.inters.TestEntity = append(c.inters.TestEntity, interceptors...)
}

// Create returns a builder for creating a TestEntity entity.
func (c *TestEntityClient) Create() *TestEntityCreate {
	mutation := newTestEntityMutation(c.config, OpCreate)
	return &TestEntityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TestEntity entities.
func (c *TestEntityClient) CreateBulk(builders ...*TestEntityCreate) *TestEntityCreateBulk {
	return &TestEntityCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TestEntityClient) MapCreateBulk(slice any, setFunc func(*TestEntityCreate, int)) *TestEntityCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TestEntityCreateBulk{err: fmt.Errorf("calling to TestEntityClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TestEntityCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TestEntityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TestEntity.
func (c *TestEntityClient) Update() *TestEntityUpdate {
	mutation := newTestEntityMutation(c.config, OpUpdate)
	return &TestEntityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TestEntityClient) UpdateOne(te *TestEntity) *TestEntityUpdateOne {
	mutation := newTestEntityMutation(c.config, OpUpdateOne, withTestEntity(te))
	return &TestEntityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TestEntityClient) UpdateOneID(id int) *TestEntityUpdateOne {
	mutation := newTestEntityMutation(c.config, OpUpdateOne, withTestEntityID(id))
	return &TestEntityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TestEntity.
func (c *TestEntityClient) Delete() *TestEntityDelete {
	mutation := newTestEntityMutation(c.config, OpDelete)
	return &TestEntityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TestEntityClient) DeleteOne(te *TestEntity) *TestEntityDeleteOne {
	return c.DeleteOneID(te.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TestEntityClient) DeleteOneID(id int) *TestEntityDeleteOne {
	builder := c.Delete().Where(testentity.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TestEntityDeleteOne{builder}
}

// Query returns a query builder for TestEntity.
func (c *TestEntityClient) Query() *TestEntityQuery {
	return &TestEntityQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTestEntity},
		inters: c.Interceptors(),
	}
}

// Get returns a TestEntity entity by its id.
func (c *TestEntityClient) Get(ctx context.Context, id int) (*TestEntity, error) {
	return c.Query().Where(testentity.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TestEntityClient) GetX(ctx context.Context, id int) *TestEntity {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TestEntityClient) Hooks() []Hook {
	return c.hooks.TestEntity
}

// Interceptors returns the client interceptors.
func (c *TestEntityClient) Interceptors() []Interceptor {
	return c.inters.TestEntity
}

func (c *TestEntityClient) mutate(ctx context.Context, m *TestEntityMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TestEntityCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TestEntityUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TestEntityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TestEntityDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TestEntity mutation op: %q", m.Op())
	}
}

// TestEntityHistoryClient is a client for the TestEntityHistory schema.
type TestEntityHistoryClient struct {
	config
}

// NewTestEntityHistoryClient returns a client for the TestEntityHistory from the given config.
func NewTestEntityHistoryClient(c config) *TestEntityHistoryClient {
	return &TestEntityHistoryClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `testentityhistory.Hooks(f(g(h())))`.
func (c *TestEntityHistoryClient) Use(hooks ...Hook) {
	c.hooks.TestEntityHistory = append(c.hooks.TestEntityHistory, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `testentityhistory.Intercept(f(g(h())))`.
func (c *TestEntityHistoryClient) Intercept(interceptors ...Interceptor) {
	c.inters.TestEntityHistory = append(c.inters.TestEntityHistory, interceptors...)
}

// Create returns a builder for creating a TestEntityHistory entity.
func (c *TestEntityHistoryClient) Create() *TestEntityHistoryCreate {
	mutation := newTestEntityHistoryMutation(c.config, OpCreate)
	return &TestEntityHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TestEntityHistory entities.
func (c *TestEntityHistoryClient) CreateBulk(builders ...*TestEntityHistoryCreate) *TestEntityHistoryCreateBulk {
	return &TestEntityHistoryCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TestEntityHistoryClient) MapCreateBulk(slice any, setFunc func(*TestEntityHistoryCreate, int)) *TestEntityHistoryCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TestEntityHistoryCreateBulk{err: fmt.Errorf("calling to TestEntityHistoryClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TestEntityHistoryCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TestEntityHistoryCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TestEntityHistory.
func (c *TestEntityHistoryClient) Update() *TestEntityHistoryUpdate {
	mutation := newTestEntityHistoryMutation(c.config, OpUpdate)
	return &TestEntityHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TestEntityHistoryClient) UpdateOne(teh *TestEntityHistory) *TestEntityHistoryUpdateOne {
	mutation := newTestEntityHistoryMutation(c.config, OpUpdateOne, withTestEntityHistory(teh))
	return &TestEntityHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TestEntityHistoryClient) UpdateOneID(id int) *TestEntityHistoryUpdateOne {
	mutation := newTestEntityHistoryMutation(c.config, OpUpdateOne, withTestEntityHistoryID(id))
	return &TestEntityHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TestEntityHistory.
func (c *TestEntityHistoryClient) Delete() *TestEntityHistoryDelete {
	mutation := newTestEntityHistoryMutation(c.config, OpDelete)
	return &TestEntityHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TestEntityHistoryClient) DeleteOne(teh *TestEntityHistory) *TestEntityHistoryDeleteOne {
	return c.DeleteOneID(teh.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TestEntityHistoryClient) DeleteOneID(id int) *TestEntityHistoryDeleteOne {
	builder := c.Delete().Where(testentityhistory.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TestEntityHistoryDeleteOne{builder}
}

// Query returns a query builder for TestEntityHistory.
func (c *TestEntityHistoryClient) Query() *TestEntityHistoryQuery {
	return &TestEntityHistoryQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTestEntityHistory},
		inters: c.Interceptors(),
	}
}

// Get returns a TestEntityHistory entity by its id.
func (c *TestEntityHistoryClient) Get(ctx context.Context, id int) (*TestEntityHistory, error) {
	return c.Query().Where(testentityhistory.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TestEntityHistoryClient) GetX(ctx context.Context, id int) *TestEntityHistory {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TestEntityHistoryClient) Hooks() []Hook {
	return c.hooks.TestEntityHistory
}

// Interceptors returns the client interceptors.
func (c *TestEntityHistoryClient) Interceptors() []Interceptor {
	return c.inters.TestEntityHistory
}

func (c *TestEntityHistoryClient) mutate(ctx context.Context, m *TestEntityHistoryMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TestEntityHistoryCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TestEntityHistoryUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TestEntityHistoryUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TestEntityHistoryDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TestEntityHistory mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		TestEntity, TestEntityHistory []ent.Hook
	}
	inters struct {
		TestEntity, TestEntityHistory []ent.Interceptor
	}
)
