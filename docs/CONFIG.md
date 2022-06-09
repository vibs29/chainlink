[//]: # (Documentation generated from docs.toml - DO NOT EDIT.)

## Table of contents

- [Global](#Global)
- [Feature](#Feature)
- [Database](#Database)
	- [Backup](#Database-Backup)
	- [Listener](#Database-Listener)
	- [Lock](#Database-Lock)
- [TelemetryIngress](#TelemetryIngress)
- [Log](#Log)
- [WebServer](#WebServer)
	- [RateLimit](#WebServer-RateLimit)
	- [MFA](#WebServer-MFA)
	- [TLS](#WebServer-TLS)
- [JobPipeline](#JobPipeline)
- [FluxMonitor](#FluxMonitor)
- [OCR2](#OCR2)
- [OCR](#OCR)
- [P2P](#P2P)
	- [V1](#P2P-V1)
	- [V2](#P2P-V2)
- [Keeper](#Keeper)
- [AutoPprof](#AutoPprof)
- [Sentry](#Sentry)
- [EVM](#EVM)
	- [BlockHistoryEstimator](#EVM-BlockHistoryEstimator)
	- [HeadTracker](#EVM-HeadTracker)
	- [KeySpecific](#EVM-KeySpecific)
	- [NodePool](#EVM-NodePool)
	- [Nodes](#EVM-Nodes)
- [Solana](#Solana)
	- [Nodes](#Solana-Nodes)
- [Terra](#Terra)
	- [Nodes](#Terra-Nodes)

## Global<a id='Global'></a>
```toml
Dev = false # Default
ExplorerURL = 'http://explorer.url' # Example
InsecureFastScrypt = false # Default
ReaperExpiration = '240h' # Default
RootDir = '~/.chainlink' # Default
ShutdownGracePeriod = '5s' # Default
```


TODO check changelog for notes on undocumented fields

### Dev<a id='Dev'></a>
```toml
Dev = false # Default
```
Setting to `true` enables development mode. This setting is not recommended for production deployments. It can be useful for enabling experimental features and collecting debug information.

### ExplorerURL<a id='ExplorerURL'></a>
```toml
ExplorerURL = 'http://explorer.url' # Example
```
The Explorer websocket URL for the node to push stats to.

### InsecureFastScrypt<a id='InsecureFastScrypt'></a>
```toml
InsecureFastScrypt = false # Default
```
**ADVANCED**

InsecureFastScrypt causes all key stores to encrypt using "fast" scrypt params instead. This is insecure and only useful for local testing. DO NOT ENABLE THIS IN PRODUCTION.

### ReaperExpiration<a id='ReaperExpiration'></a>
```toml
ReaperExpiration = '240h' # Default
```
ReaperExpiration represents how long an API session lasts before expiring and requiring a new login.

### RootDir<a id='RootDir'></a>
```toml
RootDir = '~/.chainlink' # Default
```
The Chainlink node's root directory. This is the default directory for logging, database backups, cookies, and other misc Chainlink node files. Chainlink nodes will always ensure this directory has 700 permissions because it might contain sensitive data.

### ShutdownGracePeriod<a id='ShutdownGracePeriod'></a>
```toml
ShutdownGracePeriod = '5s' # Default
```


## Feature<a id='Feature'></a>
```toml
[Feature]
FeedsManager = false # Default
LogPoller = false # Default
OffchainReporting2 = false # Default
OffchainReporting = false # Default
```


### FeedsManager<a id='Feature-FeedsManager'></a>
```toml
FeedsManager = false # Default
```
FeedsManager enables the experimental feeds manager service.

### LogPoller<a id='Feature-LogPoller'></a>
```toml
LogPoller = false # Default
```


### OffchainReporting2<a id='Feature-OffchainReporting2'></a>
```toml
OffchainReporting2 = false # Default
```
Set to `true` to enable OCR2 jobs.

### OffchainReporting<a id='Feature-OffchainReporting'></a>
```toml
OffchainReporting = false # Default
```
Set to `true` to enable OCR jobs.

## Database<a id='Database'></a>
```toml
[Database]
DefaultIdleInTxSessionTimeout = '1m'
DefaultLockTimeout = '1h'
DefaultQueryTimeout = '1s'
MigrateOnStartup = true # Default
ORMMaxIdleConns = 10 # Default
ORMMaxOpenConns = 20 # Default
```


### DefaultIdleInTxSessionTimeout<a id='Database-DefaultIdleInTxSessionTimeout'></a>
```toml
DefaultIdleInTxSessionTimeout = '1m'
```


### DefaultLockTimeout<a id='Database-DefaultLockTimeout'></a>
```toml
DefaultLockTimeout = '1h'
```


### DefaultQueryTimeout<a id='Database-DefaultQueryTimeout'></a>
```toml
DefaultQueryTimeout = '1s'
```


### MigrateOnStartup<a id='Database-MigrateOnStartup'></a>
```toml
MigrateOnStartup = true # Default
```
This variable controls whether a Chainlink node will attempt to automatically migrate the database on boot. If you want more control over your database migration process, set this variable to `false` and manually migrate the database using the CLI `migrate` command instead.

### ORMMaxIdleConns<a id='Database-ORMMaxIdleConns'></a>
```toml
ORMMaxIdleConns = 10 # Default
```
This setting configures the maximum number of idle database connections that the Chainlink node will keep open. Think of this as the baseline number of database connections per Chainlink node instance. Increasing this number can help to improve performance under database-heavy workloads.

Postgres has connection limits, so you must use cation when increasing this value. If you are running several instances of a Chainlink node or another application on a single database server, you might run out of Postgres connection slots if you raise this value too high.

### ORMMaxOpenConns<a id='Database-ORMMaxOpenConns'></a>
```toml
ORMMaxOpenConns = 20 # Default
```
This setting configures the maximum number of database connections that a Chainlink node will have open at any one time. Think of this as the maximum burst upper bound limit of database connections per Chainlink node instance. Increasing this number can help to improve performance under database-heavy workloads.

Postgres has connection limits, so you must use cation when increasing this value. If you are running several instances of a Chainlink node or another application on a single database server, you might run out of Postgres connection slots if you raise this value too high.

## Database.Backup<a id='Database-Backup'></a>
```toml
[Database.Backup]
Mode = 'none' # Default
Dir = 'test/backup/dir' # Example
OnVersionUpgrade = true
URL = 'http://test.back.up/fake'
Frequency = '1h' # Default
```
As a best practice, take regular database backups in case of accidental data loss. This best practice is especially important when you upgrade your Chainlink node to a new version. Chainlink nodes support automated database backups to make this process easier.

NOTE: Dumps can cause high load and massive database latencies, which will negatively impact the normal functioning of the Chainlink node. For this reason, it is recommended to set a `URL` and point it to a read replica if you enable automatic backups.

### Mode<a id='Database-Backup-Mode'></a>
```toml
Mode = 'none' # Default
```
Set the mode for automatic database backups, which can be one of _none_, `lite`, or `full`. If enabled, the Chainlink node will always dump a backup on every boot before running migrations. Additionally, it will automatically take database backups that overwrite the backup file for the given version at regular intervals if `Frequency` is set to a non-zero interval.

_none_ - Disables backups.
`lite` - Dumps small tables including configuration and keys that are essential for the node to function, which excludes historical data like job runs, transaction history, etc.
`full` - Dumps the entire database.

It will write to a file like `$ROOT/backup/cl_backup_<VERSION>.dump`. There is one backup dump file per version of the Chainlink node. If you upgrade the node, it will keep the backup taken right before the upgrade migration so you can restore to an older version if necessary.

### Dir<a id='Database-Backup-Dir'></a>
```toml
Dir = 'test/backup/dir' # Example
```
This variable sets the directory to use for saving the backup file. Use this if you want to save the backup file in a directory other than the default ROOT directory.

### OnVersionUpgrade<a id='Database-Backup-OnVersionUpgrade'></a>
```toml
OnVersionUpgrade = true
```


### URL<a id='Database-Backup-URL'></a>
```toml
URL = 'http://test.back.up/fake'
```
If specified, the automatic database backup will pull from this URL rather than the main `DATABASE_URL`. It is recommended to set this value to a read replica if you have one to avoid excessive load on the main database.

### Frequency<a id='Database-Backup-Frequency'></a>
```toml
Frequency = '1h' # Default
```
If this variable is set to a positive duration and `Mode` is not _none_, the node will dump the database at this regular interval.

Set to `0` to disable periodic backups.

## Database.Listener<a id='Database-Listener'></a>
```toml
[Database.Listener]
MaxReconnectDuration = '10m' # Default
MinReconnectInterval = '1m' # Default
TriggerFallbackDBPollInterval = '30s' # Default
```
**ADVANCED**

These settings control the postgres event listener.

### MaxReconnectDuration<a id='Database-Listener-MaxReconnectDuration'></a>
```toml
MaxReconnectDuration = '10m' # Default
```
MaxReconnectInterval is the maximum duration to wait between reconnect attempts.

### MinReconnectInterval<a id='Database-Listener-MinReconnectInterval'></a>
```toml
MinReconnectInterval = '1m' # Default
```
MinReconnectInterval controls the duration to wait before trying to re-establish the database connection after connection loss. After each consecutive failure this interval is doubled, until MaxReconnectInterval is reached.  Successfully completing the connection establishment procedure resets the interval back to MinReconnectInterval.

### TriggerFallbackDBPollInterval<a id='Database-Listener-TriggerFallbackDBPollInterval'></a>
```toml
TriggerFallbackDBPollInterval = '30s' # Default
```


## Database.Lock<a id='Database-Lock'></a>
```toml
[Database.Lock]
Mode = 'dual' # Default
AdvisoryCheckInterval = '1s' # Default
AdvisoryID = 1027321974924625846 # Default
LeaseDuration = '10s' # Default
LeaseRefreshInterval = '1s' # Default
```
**ADVANCED**

Do not change these settings unless you know what you are doing.

Ideally, you should use a container orchestration system like [Kubernetes](https://kubernetes.io/) to ensure that only one Chainlink node instance can ever use a specific Postgres database. However, some node operators do not have the technical capacity to do this. Common use cases run multiple Chainlink node instances in failover mode as recommended by our official documentation. The first instance takes a lock on the database and subsequent instances will wait trying to take this lock in case the first instance fails.

By default, Chainlink nodes use the `dual` setting to provide both advisory locks and lease locks for backward and forward compatibility. Using advisory locks alone presents the following problems:

- If your nodes or applications hold locks open for several hours or days, Postgres is unable to complete internal cleanup tasks. The Postgres maintainers explicitly discourage holding locks open for long periods of time.
- Advisory locks can silently disappear when you upgrade Postgres, so a new Chainlink node instance can take over even while the old node is still running.
- Advisory locks do not work well with pooling tools such as [pgbouncer](https://www.pgbouncer.org/).
- If the Chainlink node crashes, an advisory lock can hang around for up to several hours, which might require you to manually remove it so another instance of the Chainlink node will allow itself to boot.

Because of the complications with advisory locks, Chainlink nodes with v1.1.0 and later support a new `lease` locking mode. This mode might become the default in future. The `lease` locking mode works using the following process:

- Node A creates one row in the database with the client ID and updates it once per second.
- Node B spinlocks and checks periodically to see if the client ID is too old. If the client ID is not updated after a period of time, node B assumes that node A failed and takes over. Node B becomes the owner of the row and updates the client ID once per second.
- If node A comes back, it attempts to take out a lease, realizes that the database has been leased to another process, and exits the entire application immediately.

### Mode<a id='Database-Lock-Mode'></a>
```toml
Mode = 'dual' # Default
```
The Mode variable can be set to 'dual', 'advisorylock', 'lease', or 'none'. It controls which mode to use to enforce that only one Chainlink node can use the database. It is recommended to set this to `lease`.

- `dual` - The default: Uses both advisory locks and lease locks for backward and forward compatibility
- `advisorylock` - Advisory lock only
- `lease` - Lease lock only
- _none_ - No locking at all: This option useful for advanced deployment environments when you are sure that only one instance of a Chainlink node will ever be running.

### AdvisoryCheckInterval<a id='Database-Lock-AdvisoryCheckInterval'></a>
```toml
AdvisoryCheckInterval = '1s' # Default
```
This setting applies only if `Mode` is set to enable advisory locking.

Controls how often the Chainlink node checks to make sure it still holds the advisory lock when advisory locking is enabled. If a node no longer holds the lock, it will try to re-acquire it. If the node cannot re-acquire the lock, the application will exit.

### AdvisoryID<a id='Database-Lock-AdvisoryID'></a>
```toml
AdvisoryID = 1027321974924625846 # Default
```
This setting applies only if `Mode` is set to enable advisory locking.

The application advisory lock ID must match all other Chainlink nodes that might access this database. It is unlikely you will ever need to change this from the default.

### LeaseDuration<a id='Database-Lock-LeaseDuration'></a>
```toml
LeaseDuration = '10s' # Default
```
This setting applies only if `Mode` is set to enable lease locking.

How long the lease lock will last before expiring.

### LeaseRefreshInterval<a id='Database-Lock-LeaseRefreshInterval'></a>
```toml
LeaseRefreshInterval = '1s' # Default
```
This setting applies only if Mode is set to enable lease locking.

How often to refresh the lease lock. Also controls how often a standby node will check to see if it can grab the lease.

## TelemetryIngress<a id='TelemetryIngress'></a>
```toml
[TelemetryIngress]
UniConn = true # Default
Logging = false # Default
ServerPubKey = 'test-pub-key' # Default
URL = 'https://prom.test' # Example
BufferSize = 100 # Default
MaxBatchSize = 50 # Default
SendInterval = '500ms' # Default
SendTimeout = '10s' # Default
UseBatchSend = true # Default
```


### UniConn<a id='TelemetryIngress-UniConn'></a>
```toml
UniConn = true # Default
```
Toggles which ws connection style is used.

### Logging<a id='TelemetryIngress-Logging'></a>
```toml
Logging = false # Default
```
Toggles verbose logging of the raw telemetry messages being sent.

### ServerPubKey<a id='TelemetryIngress-ServerPubKey'></a>
```toml
ServerPubKey = 'test-pub-key' # Default
```
The public key of the telemetry server.

### URL<a id='TelemetryIngress-URL'></a>
```toml
URL = 'https://prom.test' # Example
```
The URL to connect to for sending telemetry.

### BufferSize<a id='TelemetryIngress-BufferSize'></a>
```toml
BufferSize = 100 # Default
```
The number of telemetry messages to buffer before dropping new ones.

### MaxBatchSize<a id='TelemetryIngress-MaxBatchSize'></a>
```toml
MaxBatchSize = 50 # Default
```
The maximum number of messages to batch into one telemetry request.

### SendInterval<a id='TelemetryIngress-SendInterval'></a>
```toml
SendInterval = '500ms' # Default
```
The interval on which batched telemetry is sent to the ingress server.

### SendTimeout<a id='TelemetryIngress-SendTimeout'></a>
```toml
SendTimeout = '10s' # Default
```
The max duration to wait for the request to complete when sending batch telemetry.

### UseBatchSend<a id='TelemetryIngress-UseBatchSend'></a>
```toml
UseBatchSend = true # Default
```
Toggles sending telemetry to the ingress server using the batch client.

## Log<a id='Log'></a>
```toml
[Log]
DatabaseQueries = false # Default
JSONConsole = false # Default
FileDir = '/my/log/directory' # Example
FileMaxSize = '5120mb'
FileMaxAgeDays = 0 # Default
FileMaxBackups = 1 # Default
UnixTS = false # Default
```


### DatabaseQueries<a id='Log-DatabaseQueries'></a>
```toml
DatabaseQueries = false # Default
```
This setting tells the Chainlink node to log database queries made using the default logger. SQL statements will be logged at `debug` level. Not all statements can be logged. The best way to get a true log of all SQL statements is to enable SQL statement logging on Postgres.

### JSONConsole<a id='Log-JSONConsole'></a>
```toml
JSONConsole = false # Default
```
Set this to true to enable JSON logging. Otherwise, the log is saved in a human-friendly console format.

### FileDir<a id='Log-FileDir'></a>
```toml
FileDir = '/my/log/directory' # Example
```
By default, Chainlink nodes write log data to `$ROOT/log.jsonl`. The log directory can be changed by setting this var.

### FileMaxSize<a id='Log-FileMaxSize'></a>
```toml
FileMaxSize = '5120mb'
```
Determines the log file's max size in megabytes before file rotation. Having this not set will disable logging to disk. If your disk doesn't have enough disk space, the logging will pause and the application will log errors until space is available again.

Values must have suffixes with a unit like: `5120mb` (5,120 megabytes). If no unit suffix is provided, the value defaults to `b` (bytes). The list of valid unit suffixes are:

- b (bytes)
- kb (kilobytes)
- mb (megabytes)
- gb (gigabytes)
- tb (terabytes)

### FileMaxAgeDays<a id='Log-FileMaxAgeDays'></a>
```toml
FileMaxAgeDays = 0 # Default
```
Determines the log file's max age in days before file rotation. Keeping this config with the default value will not remove log files based on age.

### FileMaxBackups<a id='Log-FileMaxBackups'></a>
```toml
FileMaxBackups = 1 # Default
```
Determines the maximum number of old log files to retain. Keeping this config with the default value retains all old log files. The `FileMaxAgeDays` variable can still cause them to get deleted.

### UnixTS<a id='Log-UnixTS'></a>
```toml
UnixTS = false # Default
```
Previous versions of Chainlink nodes wrote JSON logs with a unix timestamp. As of v1.1.0 and up, the default has changed to use ISO8601 timestamps for better readability. Setting `true` will enable the old behavior.

## WebServer<a id='WebServer'></a>
```toml
[WebServer]
AllowOrigins = 'http://localhost:3000,http://localhost:6688' # Default
BridgeResponseURL = 'https://my-chainlink-node.example.com:6688' # Example
HTTPWriteTimeout = '10s' # Default
HTTPPort = 6688 # Default
SecureCookies = true # Default
SessionTimeout = '15m' # Default
```


### AllowOrigins<a id='WebServer-AllowOrigins'></a>
```toml
AllowOrigins = 'http://localhost:3000,http://localhost:6688' # Default
```
Controls the URLs Chainlink nodes emit in the `Allow-Origins` header of its API responses. The setting can be a comma-separated list with no spaces. You might experience CORS issues if this is not set correctly.

You should set this to the external URL that you use to access the Chainlink UI.

You can set `AllowOrigins = '*'` to allow the UI to work from any URL, but it is recommended for security reasons to make it explicit instead.

### BridgeResponseURL<a id='WebServer-BridgeResponseURL'></a>
```toml
BridgeResponseURL = 'https://my-chainlink-node.example.com:6688' # Example
```
Defines the URL for bridges to send a response to. This _must_ be set when using async external adapters.

Usually this will be the same as the URL/IP and port you use to connect to the Chainlink UI.

### HTTPWriteTimeout<a id='WebServer-HTTPWriteTimeout'></a>
```toml
HTTPWriteTimeout = '10s' # Default
```
**ADVANCED**

Do not change this setting unless you know what you are doing.

Controls how long the Chainlink node's API server can hold a socket open for writing a response to an HTTP request. Sometimes, this must be increased for pprof.

### HTTPPort<a id='WebServer-HTTPPort'></a>
```toml
HTTPPort = 6688 # Default
```
Port used for the Chainlink Node API, [CLI](/docs/configuration-variables/#cli-client), and GUI.

### SecureCookies<a id='WebServer-SecureCookies'></a>
```toml
SecureCookies = true # Default
```
Requires the use of secure cookies for authentication. Set to false to enable standard HTTP requests along with `TLSPort = 0`.

### SessionTimeout<a id='WebServer-SessionTimeout'></a>
```toml
SessionTimeout = '15m' # Default
```
This value determines the amount of idle time to elapse before session cookies expire. This signs out GUI users from their sessions.

## WebServer.RateLimit<a id='WebServer-RateLimit'></a>
```toml
[WebServer.RateLimit]
Authenticated = 42 # Default
AuthenticatedPeriod = '1m' # Default
Unauthenticated = 5 # Default
UnauthenticatedPeriod = '20s' # Default
```


### Authenticated<a id='WebServer-RateLimit-Authenticated'></a>
```toml
Authenticated = 42 # Default
```
Defines the threshold to which authenticated requests get limited. More than this many authenticated requests per `AuthenticatedRateLimitPeriod` will be rejected.

### AuthenticatedPeriod<a id='WebServer-RateLimit-AuthenticatedPeriod'></a>
```toml
AuthenticatedPeriod = '1m' # Default
```
Defines the period to which authenticated requests get limited.

### Unauthenticated<a id='WebServer-RateLimit-Unauthenticated'></a>
```toml
Unauthenticated = 5 # Default
```
Defines the threshold to which authenticated requests get limited. More than this many unauthenticated requests per `UnAuthenticatedRateLimitPeriod` will be rejected.

### UnauthenticatedPeriod<a id='WebServer-RateLimit-UnauthenticatedPeriod'></a>
```toml
UnauthenticatedPeriod = '20s' # Default
```
Defines the period to which unauthenticated requests get limited.

## WebServer.MFA<a id='WebServer-MFA'></a>
```toml
[WebServer.MFA]
RPID = 'localhost' # Example
RPOrigin = 'http://localhost:6688/' # Example
```
The Operator UI frontend supports enabling Multi Factor Authentication via Webauthn per account. When enabled, logging in will require the account password and a hardware or OS security key such as Yubikey. To enroll, log in to the operator UI and click the circle purple profile button at the top right and then click **Register MFA Token**. Tap your hardware security key or use the OS public key management feature to enroll a key. Next time you log in, this key will be required to authenticate.

### RPID<a id='WebServer-MFA-RPID'></a>
```toml
RPID = 'localhost' # Example
```
The FQDN of where the Operator UI is served. When serving locally, the value should be `localhost`.

### RPOrigin<a id='WebServer-MFA-RPOrigin'></a>
```toml
RPOrigin = 'http://localhost:6688/' # Example
```
The origin URL where WebAuthn requests initiate, including scheme and port. When serving locally, the value should be `http://localhost:6688/`.

## WebServer.TLS<a id='WebServer-TLS'></a>
```toml
[WebServer.TLS]
CertPath = '/home/$USER/.chainlink/tls/server.crt' # Example
Host = 'tls-host' # Example
KeyPath = '/home/$USER/.chainlink/tls/server.key' # Example
HTTPSPort = 6689 # Default
ForceRedirect = false # Default
```
The TLS settings apply only if you want to enable TLS security on your Chainlink node.

### CertPath<a id='WebServer-TLS-CertPath'></a>
```toml
CertPath = '/home/$USER/.chainlink/tls/server.crt' # Example
```
The location of the TLS certificate file.

### Host<a id='WebServer-TLS-Host'></a>
```toml
Host = 'tls-host' # Example
```
The hostname configured for TLS to be used by the Chainlink node. This is useful if you configured a domain name specific for your Chainlink node.

### KeyPath<a id='WebServer-TLS-KeyPath'></a>
```toml
KeyPath = '/home/$USER/.chainlink/tls/server.key' # Example
```
The location of the TLS private key file.

### HTTPSPort<a id='WebServer-TLS-HTTPSPort'></a>
```toml
HTTPSPort = 6689 # Default
```
The port used for HTTPS connections. Set this to `0` to disable HTTPS. Disabling HTTPS also relieves Chainlink nodes of the requirement for a TLS certificate.

### ForceRedirect<a id='WebServer-TLS-ForceRedirect'></a>
```toml
ForceRedirect = false # Default
```
Forces TLS redirect for unencrypted connections.

## JobPipeline<a id='JobPipeline'></a>
```toml
[JobPipeline]
HTTPRequestMaxSize = '32768' # Default
DefaultHTTPRequestTimeout = '15s' # Default
ExternalInitiatorsEnabled = false # Default
MaxRunDuration = '10m' # Default
ReaperInterval = '1h' # Default
ReaperThreshold = '24h' # Default
ResultWriteQueueDepth = 100 # Default
```


### HTTPRequestMaxSize<a id='JobPipeline-HTTPRequestMaxSize'></a>
```toml
HTTPRequestMaxSize = '32768' # Default
```
Defines the maximum size for HTTP requests and responses made by `http` and `bridge` adapters.

### DefaultHTTPRequestTimeout<a id='JobPipeline-DefaultHTTPRequestTimeout'></a>
```toml
DefaultHTTPRequestTimeout = '15s' # Default
```
Defines the default timeout for HTTP requests made by `http` and `bridge` adapters.

### ExternalInitiatorsEnabled<a id='JobPipeline-ExternalInitiatorsEnabled'></a>
```toml
ExternalInitiatorsEnabled = false # Default
```
Enables the External Initiator feature. If disabled, `webhook` jobs can ONLY be initiated by a logged-in user. If enabled, `webhook` jobs can be initiated by a whitelisted external initiator.

### MaxRunDuration<a id='JobPipeline-MaxRunDuration'></a>
```toml
MaxRunDuration = '10m' # Default
```
The maximum time allowed for a single job run. If it takes longer, it will exit early and be marked errored. If set to zero, disables the time limit completely.

### ReaperInterval<a id='JobPipeline-ReaperInterval'></a>
```toml
ReaperInterval = '1h' # Default
```
In order to keep database size manageable, Chainlink nodes will run a reaper that deletes completed job runs older than a certain threshold age. `ReaperInterval` controls how often the job pipeline reaper will run.

Set to `0` to disable the periodic reaper.

### ReaperThreshold<a id='JobPipeline-ReaperThreshold'></a>
```toml
ReaperThreshold = '24h' # Default
```
Determines the age limit for job runs. Completed job runs older than this will be automatically purged from the database.

### ResultWriteQueueDepth<a id='JobPipeline-ResultWriteQueueDepth'></a>
```toml
ResultWriteQueueDepth = 100 # Default
```
Some jobs write their results asynchronously for performance reasons such as OCR. `JOB_PIPELINE_RESULT_WRITE_QUEUE_DEPTH` controls how many writes will be buffered before subsequent writes are dropped.

Do not change this setting unless you know what you are doing.

## FluxMonitor<a id='FluxMonitor'></a>
```toml
[FluxMonitor]
DefaultTransactionQueueDepth = 1 # Default
SimulateTransactions = false # Default
```


### DefaultTransactionQueueDepth<a id='FluxMonitor-DefaultTransactionQueueDepth'></a>
```toml
DefaultTransactionQueueDepth = 1 # Default
```
**ADVANCED**

DefaultTransactionQueueDepth controls the queue size for `DropOldestStrategy` in Flux Monitor. Set to 0 to use `SendEvery` strategy instead.

### SimulateTransactions<a id='FluxMonitor-SimulateTransactions'></a>
```toml
SimulateTransactions = false # Default
```
Enable transaction simulation for Flux Monitor.

## OCR2<a id='OCR2'></a>
```toml
[OCR2]
ContractConfirmations = 3 # Default
BlockchainTimeout = '20s' # Default
ContractPollInterval = '1m' # Default
ContractSubscribeInterval = '2m' # Default
ContractTransmitterTransmitTimeout = '10s' # Default
DatabaseTimeout = '10s' # Default
KeyBundleID = '7a5f66bbe6594259325bf2b4f5b1a9c900000000000000000000000000000000' # Example
MonitoringEndpoint = 'test-mon-end' # Example
```


### ContractConfirmations<a id='OCR2-ContractConfirmations'></a>
```toml
ContractConfirmations = 3 # Default
```


### BlockchainTimeout<a id='OCR2-BlockchainTimeout'></a>
```toml
BlockchainTimeout = '20s' # Default
```


### ContractPollInterval<a id='OCR2-ContractPollInterval'></a>
```toml
ContractPollInterval = '1m' # Default
```


### ContractSubscribeInterval<a id='OCR2-ContractSubscribeInterval'></a>
```toml
ContractSubscribeInterval = '2m' # Default
```


### ContractTransmitterTransmitTimeout<a id='OCR2-ContractTransmitterTransmitTimeout'></a>
```toml
ContractTransmitterTransmitTimeout = '10s' # Default
```


### DatabaseTimeout<a id='OCR2-DatabaseTimeout'></a>
```toml
DatabaseTimeout = '10s' # Default
```


### KeyBundleID<a id='OCR2-KeyBundleID'></a>
```toml
KeyBundleID = '7a5f66bbe6594259325bf2b4f5b1a9c900000000000000000000000000000000' # Example
```
TODO

### MonitoringEndpoint<a id='OCR2-MonitoringEndpoint'></a>
```toml
MonitoringEndpoint = 'test-mon-end' # Example
```
TODO

## OCR<a id='OCR'></a>
```toml
[OCR]
ObservationTimeout = '11s'
BlockchainTimeout = '3s'
ContractPollInterval = '1h'
ContractSubscribeInterval = '1m'
DefaultTransactionQueueDepth = 12
KeyBundleID = 'acdd42797a8b921b2910497badc5000600000000000000000000000000000000' # Example
MonitoringEndpoint = 'test-monitor' # Example
SimulateTransactions = false # Default
TransmitterAddress = '0xa0788FC17B1dEe36f057c42B6F373A34B014687e' # Example
TraceLogging = false # Default
```
This section applies only if you are running off-chain reporting jobs.

### ObservationTimeout<a id='OCR-ObservationTimeout'></a>
```toml
ObservationTimeout = '11s'
```


### BlockchainTimeout<a id='OCR-BlockchainTimeout'></a>
```toml
BlockchainTimeout = '3s'
```


### ContractPollInterval<a id='OCR-ContractPollInterval'></a>
```toml
ContractPollInterval = '1h'
```


### ContractSubscribeInterval<a id='OCR-ContractSubscribeInterval'></a>
```toml
ContractSubscribeInterval = '1m'
```


### DefaultTransactionQueueDepth<a id='OCR-DefaultTransactionQueueDepth'></a>
```toml
DefaultTransactionQueueDepth = 12
```
**ADVANCED**

DefaultTransactionQueueDepth controls the queue size for `DropOldestStrategy` in OCR. Set to 0 to use `SendEvery` strategy instead.

### KeyBundleID<a id='OCR-KeyBundleID'></a>
```toml
KeyBundleID = 'acdd42797a8b921b2910497badc5000600000000000000000000000000000000' # Example
```
Default key bundle ID to use for OCR jobs. If you have an OCR job that does not explicitly specify a key bundle ID, it will fall back to this value.

### MonitoringEndpoint<a id='OCR-MonitoringEndpoint'></a>
```toml
MonitoringEndpoint = 'test-monitor' # Example
```
Optional URL of OCR monitoring endpoint.

### SimulateTransactions<a id='OCR-SimulateTransactions'></a>
```toml
SimulateTransactions = false # Default
```
Enable transaction simulation for OCR.

### TransmitterAddress<a id='OCR-TransmitterAddress'></a>
```toml
TransmitterAddress = '0xa0788FC17B1dEe36f057c42B6F373A34B014687e' # Example
```
The default sending address to use for OCR. If you have an OCR job that does not explicitly specify a transmitter address, it will fall back to this value.

### TraceLogging<a id='OCR-TraceLogging'></a>
```toml
TraceLogging = false # Default
```


## P2P<a id='P2P'></a>
```toml
[P2P]
IncomingMessageBufferSize = 10 # Default
OutgoingMessageBufferSize = 10 # Default
```
P2P supports multiple networking stack versions. You may configure `[P2P.V1]`, `[P2P.V2]`, or both to run simultaneously.
If both are configured, then for each link with another peer, V2 networking will be preferred. If V2 does not work, the link will
automatically fall back to V1. If V2 starts working again later, it will automatically be preferred again. This is useful
for migrating networks without downtime. Note that the two networking stacks _must not_ be configured to bind to the same IP/port.

All nodes in the OCR network should share the same networking stack.

### IncomingMessageBufferSize<a id='P2P-IncomingMessageBufferSize'></a>
```toml
IncomingMessageBufferSize = 10 # Default
```


### OutgoingMessageBufferSize<a id='P2P-OutgoingMessageBufferSize'></a>
```toml
OutgoingMessageBufferSize = 10 # Default
```


## P2P.V1<a id='P2P-V1'></a>
```toml
[P2P.V1]
AnnounceIP = '1.2.3.4' # Example
AnnouncePort = 1337 # Example
BootstrapCheckInterval = '20s' # Default
DefaultBootstrapPeers = ['/dns4/example.com/tcp/1337/p2p/12D3KooWMHMRLQkgPbFSYHwD3NBuwtS1AmxhvKVUrcfyaGDASR4U', '/ip4/1.2.3.4/tcp/9999/p2p/12D3KooWLZ9uTC3MrvKfDpGju6RAQubiMDL7CuJcAgDRTYP7fh7R'] # Example
DHTAnnouncementCounterUserPrefix = 0 # Default
DHTLookupInterval = 10 # Default
ListenIP = '0.0.0.0' # Default
ListenPort = 1337 # Example
NewStreamTimeout = '10s' # Default
PeerID = '12D3KooWMoejJznyDuEk5aX6GvbjaG12UzeornPCBNzMRqdwrFJw' # Example
PeerstoreWriteInterval = '5m' # Default
```


### AnnounceIP<a id='P2P-V1-AnnounceIP'></a>
```toml
AnnounceIP = '1.2.3.4' # Example
```
Should be set as the externally reachable IP address of the Chainlink node.

### AnnouncePort<a id='P2P-V1-AnnouncePort'></a>
```toml
AnnouncePort = 1337 # Example
```
Should be set as the externally reachable port of the Chainlink node.

### BootstrapCheckInterval<a id='P2P-V1-BootstrapCheckInterval'></a>
```toml
BootstrapCheckInterval = '20s' # Default
```


### DefaultBootstrapPeers<a id='P2P-V1-DefaultBootstrapPeers'></a>
```toml
DefaultBootstrapPeers = ['/dns4/example.com/tcp/1337/p2p/12D3KooWMHMRLQkgPbFSYHwD3NBuwtS1AmxhvKVUrcfyaGDASR4U', '/ip4/1.2.3.4/tcp/9999/p2p/12D3KooWLZ9uTC3MrvKfDpGju6RAQubiMDL7CuJcAgDRTYP7fh7R'] # Example
```
Default set of bootstrap peers.

### DHTAnnouncementCounterUserPrefix<a id='P2P-V1-DHTAnnouncementCounterUserPrefix'></a>
```toml
DHTAnnouncementCounterUserPrefix = 0 # Default
```


### DHTLookupInterval<a id='P2P-V1-DHTLookupInterval'></a>
```toml
DHTLookupInterval = 10 # Default
```


### ListenIP<a id='P2P-V1-ListenIP'></a>
```toml
ListenIP = '0.0.0.0' # Default
```
The default IP address to bind to.

### ListenPort<a id='P2P-V1-ListenPort'></a>
```toml
ListenPort = 1337 # Example
```
The port to listen on. If left blank, the node randomly selects a different port each time it boots. It is highly recommended to set this to a static value to avoid network instability.

### NewStreamTimeout<a id='P2P-V1-NewStreamTimeout'></a>
```toml
NewStreamTimeout = '10s' # Default
```


### PeerID<a id='P2P-V1-PeerID'></a>
```toml
PeerID = '12D3KooWMoejJznyDuEk5aX6GvbjaG12UzeornPCBNzMRqdwrFJw' # Example
```
The default peer ID to use for OCR jobs. If unspecified, uses the first available peer ID.

### PeerstoreWriteInterval<a id='P2P-V1-PeerstoreWriteInterval'></a>
```toml
PeerstoreWriteInterval = '5m' # Default
```
**ADVANCED**

PeerstoreWriteInterval controls how often the peerstore for the OCR V1 networking stack is persisted to the database.

## P2P.V2<a id='P2P-V2'></a>
```toml
[P2P.V2]
AnnounceAddresses = ['1.2.3.4:9999', '[a52d:0:a88:1274::abcd]:1337'] # Example
DefaultBootstrappers = ['12D3KooWMHMRLQkgPbFSYHwD3NBuwtS1AmxhvKVUrcfyaGDASR4U@1.2.3.4:9999', '12D3KooWM55u5Swtpw9r8aFLQHEtw7HR4t44GdNs654ej5gRs2Dh@example.com:1234'] # Example
DeltaDial = '15s' # Default
DeltaReconcile = '1m' # Default
ListenAddresses = ['1.2.3.4:9999', '[a52d:0:a88:1274::abcd]:1337'] # Example
```


### AnnounceAddresses<a id='P2P-V2-AnnounceAddresses'></a>
```toml
AnnounceAddresses = ['1.2.3.4:9999', '[a52d:0:a88:1274::abcd]:1337'] # Example
```
The addresses the peer will advertise on the network in host:port form as accepted by net.Dial. The addresses should be reachable by peers of interest.

### DefaultBootstrappers<a id='P2P-V2-DefaultBootstrappers'></a>
```toml
DefaultBootstrappers = ['12D3KooWMHMRLQkgPbFSYHwD3NBuwtS1AmxhvKVUrcfyaGDASR4U@1.2.3.4:9999', '12D3KooWM55u5Swtpw9r8aFLQHEtw7HR4t44GdNs654ej5gRs2Dh@example.com:1234'] # Example
```
The default bootstrapper peers for libocr's v2 networking stack.

### DeltaDial<a id='P2P-V2-DeltaDial'></a>
```toml
DeltaDial = '15s' # Default
```


### DeltaReconcile<a id='P2P-V2-DeltaReconcile'></a>
```toml
DeltaReconcile = '1m' # Default
```


### ListenAddresses<a id='P2P-V2-ListenAddresses'></a>
```toml
ListenAddresses = ['1.2.3.4:9999', '[a52d:0:a88:1274::abcd]:1337'] # Example
```
The addresses the peer will listen to on the network in `host:port` form as accepted by `net.Listen()`, but the host and port must be fully specified and cannot be empty. You can specify `0.0.0.0` (IPv4) or `::` (IPv6) to listen on all interfaces, but that is not recommended.

## Keeper<a id='Keeper'></a>
```toml
[Keeper]
DefaultTransactionQueueDepth = 1 # Default
GasPriceBufferPercent = 20 # Default
GasTipCapBufferPercent = 20 # Default
BaseFeeBufferPercent = 20 # Default
MaximumGracePeriod = 100 # Default
RegistryCheckGasOverhead = '200000' # Default
RegistryPerformGasOverhead = '150000' # Default
RegistrySyncInterval = '30m' # Default
RegistrySyncUpkeepQueueSize = 10 # Default
TurnLookBack = 1000 # Default
TurnFlagEnabled = false # Default
UpkeepCheckGasPriceEnabled = false # Default
```


### DefaultTransactionQueueDepth<a id='Keeper-DefaultTransactionQueueDepth'></a>
```toml
DefaultTransactionQueueDepth = 1 # Default
```
**ADVANCED**

DefaultTransactionQueueDepth controls the queue size for `DropOldestStrategy` in Keeper. Set to 0 to use `SendEvery` strategy instead.

### GasPriceBufferPercent<a id='Keeper-GasPriceBufferPercent'></a>
```toml
GasPriceBufferPercent = 20 # Default
```
Adds the specified percentage to the gas price used for checking whether to perform an upkeep. Only applies in legacy mode (EIP-1559 off).

### GasTipCapBufferPercent<a id='Keeper-GasTipCapBufferPercent'></a>
```toml
GasTipCapBufferPercent = 20 # Default
```
Adds the specified percentage to the gas price used for checking whether to perform an upkeep. Only applies in EIP-1559 mode.

### BaseFeeBufferPercent<a id='Keeper-BaseFeeBufferPercent'></a>
```toml
BaseFeeBufferPercent = 20 # Default
```
Adds the specified percentage to the base fee used for checking whether to perform an upkeep. Applies only in EIP-1559 mode.

### MaximumGracePeriod<a id='Keeper-MaximumGracePeriod'></a>
```toml
MaximumGracePeriod = 100 # Default
```
**ADVANCED**

Do not change this setting unless you know what you are doing.

The maximum number of blocks that a keeper will wait after performing an upkeep before it resumes checking that upkeep

### RegistryCheckGasOverhead<a id='Keeper-RegistryCheckGasOverhead'></a>
```toml
RegistryCheckGasOverhead = '200000' # Default
```
**ADVANCED**

Do not change this setting unless you know what you are doing.

The amount of extra gas to provide checkUpkeep() calls to account for the gas consumed by the keeper registry.

### RegistryPerformGasOverhead<a id='Keeper-RegistryPerformGasOverhead'></a>
```toml
RegistryPerformGasOverhead = '150000' # Default
```
**ADVANCED**

Do not change this setting unless you know what you are doing.

The amount of extra gas to provide performUpkeep() calls to account for the gas consumed by the keeper registry

### RegistrySyncInterval<a id='Keeper-RegistrySyncInterval'></a>
```toml
RegistrySyncInterval = '30m' # Default
```
**ADVANCED**

Do not change this setting unless you know what you are doing.

The interval in which the RegistrySynchronizer performs a full sync of the keeper registry contract it is tracking.

### RegistrySyncUpkeepQueueSize<a id='Keeper-RegistrySyncUpkeepQueueSize'></a>
```toml
RegistrySyncUpkeepQueueSize = 10 # Default
```
**ADVANCED**

Do not change this setting unless you know what you are doing.

Represents the maximum number of upkeeps that can be synced in parallel.

### TurnLookBack<a id='Keeper-TurnLookBack'></a>
```toml
TurnLookBack = 1000 # Default
```
The number of blocks in the past to look back when getting a block for a turn.

### TurnFlagEnabled<a id='Keeper-TurnFlagEnabled'></a>
```toml
TurnFlagEnabled = false # Default
```
Enables a new algorithm for how keepers take turns.

### UpkeepCheckGasPriceEnabled<a id='Keeper-UpkeepCheckGasPriceEnabled'></a>
```toml
UpkeepCheckGasPriceEnabled = false # Default
```
**ADVANCED**

Use this setting *only* on Polygon networks.

Includes gas price in calls to `checkUpkeep()` when set to `true`.

## AutoPprof<a id='AutoPprof'></a>
```toml
[AutoPprof]
Enabled = false # Default
ProfileRoot = 'prof/root' # Example
PollInterval = '10s' # Default
GatherDuration = '10s' # Default
GatherTraceDuration = '5s' # Default
MaxProfileSize = '100mb' # Default
CPUProfileRate = 1 # Default
MemProfileRate = 1 # Default
BlockProfileRate = 1 # Default
MutexProfileFraction = 1 # Default
MemThreshold = '4gb' # Default
GoroutineThreshold = 5000 # Default
```
The Chainlink node is equipped with an internal "nurse" service that can perform automatic `pprof` profiling when the certain resource thresholds are exceeded, such as memory and goroutine count. These profiles are saved to disk to facilitate fine-grained debugging of performance-related issues. In general, if you notice that your node has begun to accumulate profiles, forward them to the Chainlink team.

To learn more about these profiles, read the [Profiling Go programs with pprof](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/) guide.

### Enabled<a id='AutoPprof-Enabled'></a>
```toml
Enabled = false # Default
```
Set to `true` to enable the automatic profiling service.

### ProfileRoot<a id='AutoPprof-ProfileRoot'></a>
```toml
ProfileRoot = 'prof/root' # Example
```
The location on disk where pprof profiles will be stored. Defaults to `RootDir`.

### PollInterval<a id='AutoPprof-PollInterval'></a>
```toml
PollInterval = '10s' # Default
```
The interval at which the node's resources are checked.

### GatherDuration<a id='AutoPprof-GatherDuration'></a>
```toml
GatherDuration = '10s' # Default
```
The duration for which profiles are gathered when profiling starts.

### GatherTraceDuration<a id='AutoPprof-GatherTraceDuration'></a>
```toml
GatherTraceDuration = '5s' # Default
```
The duration for which traces are gathered when profiling is kicked off. This is separately configurable because traces are significantly larger than other types of profiles.

### MaxProfileSize<a id='AutoPprof-MaxProfileSize'></a>
```toml
MaxProfileSize = '100mb' # Default
```
The maximum amount of disk space that profiles may consume before profiling is disabled.

### CPUProfileRate<a id='AutoPprof-CPUProfileRate'></a>
```toml
CPUProfileRate = 1 # Default
```
See https://pkg.go.dev/runtime#SetCPUProfileRate.

### MemProfileRate<a id='AutoPprof-MemProfileRate'></a>
```toml
MemProfileRate = 1 # Default
```
See https://pkg.go.dev/runtime#pkg-variables.

### BlockProfileRate<a id='AutoPprof-BlockProfileRate'></a>
```toml
BlockProfileRate = 1 # Default
```
See https://pkg.go.dev/runtime#SetBlockProfileRate.

### MutexProfileFraction<a id='AutoPprof-MutexProfileFraction'></a>
```toml
MutexProfileFraction = 1 # Default
```
See https://pkg.go.dev/runtime#SetMutexProfileFraction.

### MemThreshold<a id='AutoPprof-MemThreshold'></a>
```toml
MemThreshold = '4gb' # Default
```
The maximum amount of memory the node can actively consume before profiling begins.

### GoroutineThreshold<a id='AutoPprof-GoroutineThreshold'></a>
```toml
GoroutineThreshold = 5000 # Default
```
The maximum number of actively-running goroutines the node can spawn before profiling begins.

## Sentry<a id='Sentry'></a>
```toml
[Sentry]
Debug = false # Default
DSN = 'sentry-dsn'
Environment = 'dev'
Release = 'v1.2.3'
```


### Debug<a id='Sentry-Debug'></a>
```toml
Debug = false # Default
```


### DSN<a id='Sentry-DSN'></a>
```toml
DSN = 'sentry-dsn'
```


### Environment<a id='Sentry-Environment'></a>
```toml
Environment = 'dev'
```


### Release<a id='Sentry-Release'></a>
```toml
Release = 'v1.2.3'
```


## EVM<a id='EVM'></a>
```toml
[[EVM]]
ChainID = '1'
Enabled = true # Default
BalanceMonitorEnabled = true
BlockBackfillDepth = 100
BlockBackfillSkip = true
ChainType = 'Optimism'
EIP1559DynamicFees = false # Default
FinalityDepth = 42
FlagsContractAddress = '0xae4E781a6218A8031764928E88d457937A954fC3' # Example
GasBumpPercent = 10
GasBumpThreshold = 15
GasBumpTxDepth = 6
GasBumpWei = '100'
GasEstimatorMode = 'BlockHistory' # Default
GasFeeCapDefault = '9223372036854775807'
GasLimitDefault = '12'
GasLimitMultiplier = '1.234'
GasLimitTransfer = '100'
GasPriceDefault = '9223372036854775807'
GasTipCapDefault = '2'
GasTipCapMinimum = '1'
LinkContractAddress = '0x538aAaB4ea120b2bC2fe5D296852D948F07D849e'
LogBackfillBatchSize = 17
LogPollInterval = '1m'
MaxGasPriceWei = '281474976710655'
MaxInFlightTransactions = 19
MaxQueuedTransactions = 99
MinGasPriceWei = '13'
MinIncomingConfirmations = 13
MinimumContractPayment = '9223372036854775807'
NonceAutoSync = false # Default
OCRContractConfirmations = 11
OCRContractTransmitterTransmitTimeout = '1m'
OCRDatabaseTimeout = '1s'
OCRObservationGracePeriod = '1s'
OCRObservationTimeout = '1m'
OCR2ContractConfirmations = 7
OperatorFactoryAddress = '0xa5B85635Be42F21f94F28034B7DA440EeFF0F418'
RPCDefaultBatchSize = 17
TxReaperInterval = '1m'
TxReaperThreshold = '1m'
TxResendAfterThreshold = '1h'
UseForwarders = false # Default
```


### ChainID<a id='EVM-ChainID'></a>
```toml
ChainID = '1'
```


### Enabled<a id='EVM-Enabled'></a>
```toml
Enabled = true # Default
```


### BalanceMonitorEnabled<a id='EVM-BalanceMonitorEnabled'></a>
```toml
BalanceMonitorEnabled = true
```


### BlockBackfillDepth<a id='EVM-BlockBackfillDepth'></a>
```toml
BlockBackfillDepth = 100
```


### BlockBackfillSkip<a id='EVM-BlockBackfillSkip'></a>
```toml
BlockBackfillSkip = true
```


### ChainType<a id='EVM-ChainType'></a>
```toml
ChainType = 'Optimism'
```


### EIP1559DynamicFees<a id='EVM-EIP1559DynamicFees'></a>
```toml
EIP1559DynamicFees = false # Default
```
Forces EIP-1559 transaction mode for all chains. Enabling EIP-1559 mode can help reduce gas costs on chains that support it. This is supported only on official Ethereum mainnet and testnets. It is not recommended to enable this setting on Polygon because the EIP-1559 fee market appears to be broken on all Polygon chains and EIP-1559 transactions are less likely to be included than legacy transactions.

#### Technical details

Chainlink nodes include experimental support for submitting transactions using type 0x2 (EIP-1559) envelope.

EIP-1559 mode is enabled by default on the Ethereum Mainnet, but can be enabled on a per-chain basis or globally.

This might help to save gas on spikes. Chainlink nodes should react faster on the upleg and avoid overpaying on the downleg. It might also be possible to set `BLOCK_HISTORY_ESTIMATOR_BATCH_SIZE` to a smaller value such as 12 or even 6 because tip cap should be a more consistent indicator of inclusion time than total gas price. This would make Chainlink nodes more responsive and should reduce response time variance. Some experimentation is required to find optimum settings.

Set with caution, if you set this on a chain that does not actually support EIP-1559 your node will be broken.

In EIP-1559 mode, the total price for the transaction is the minimum of base fee + tip cap and fee cap. More information can be found on the [official EIP](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1559.md).

Chainlink's implementation of EIP-1559 works as follows:

If you are using FixedPriceEstimator:
- With gas bumping disabled, it will submit all transactions with `feecap=MaxGasPriceWei` and `tipcap=GasTipCapDefault`
- With gas bumping enabled, it will submit all transactions initially with `feecap=GasFeeCapDefault` and `tipcap=GasTipCapDefault`.

If you are using BlockHistoryEstimator (default for most chains):
- With gas bumping disabled, it will submit all transactions with `feecap=MaxGasPriceWei` and `tipcap=<calculated using past blocks>`
- With gas bumping enabled (default for most chains) it will submit all transactions initially with `feecap=current block base fee * (1.125 ^ N)` where N is configurable by setting BLOCK_HISTORY_ESTIMATOR_EIP1559_FEE_CAP_BUFFER_BLOCKS but defaults to `gas bump threshold+1` and `tipcap=<calculated using past blocks>`

Bumping works as follows:

- Increase tipcap by `max(tipcap * (1 + GasBumpPercent), tipcap + GasBumpWei)`
- Increase feecap by `max(feecap * (1 + GasBumpPercent), feecap + GasBumpWei)`

A quick note on terminology - Chainlink nodes use the same terms used internally by go-ethereum source code to describe various prices. This is not the same as the externally used terms. For reference:

- Base Fee Per Gas = BaseFeePerGas
- Max Fee Per Gas = FeeCap
- Max Priority Fee Per Gas = TipCap

In EIP-1559 mode, the following changes occur to how configuration works:

- All new transactions will be sent as type 0x2 transactions specifying a TipCap and FeeCap. Be aware that existing pending legacy transactions will continue to be gas bumped in legacy mode.
- `BlockHistoryEstimator` will apply its calculations (gas percentile etc) to the TipCap and this value will be used for new transactions (GasPrice will be ignored)
- `FixedPriceEstimator` will use `GasTipCapDefault` instead of `GasPriceDefault` for the tip cap
- `FixedPriceEstimator` will use `GasFeeCapDefault` instaed of `GasPriceDefault` for the fee cap
- `MinGasPriceWei` is ignored for new transactions and `GasTipCapMinimum` is used instead (default 0)
- `MaxGasPriceWei` still represents that absolute upper limit that Chainlink will ever spend (total) on a single tx
- `KEEPER_GAS_PRICE_BUFFER_PERCENT` is ignored in EIP-1559 mode and `KEEPER_TIP_CAP_BUFFER_PERCENT` is used instead

### FinalityDepth<a id='EVM-FinalityDepth'></a>
```toml
FinalityDepth = 42
```


### FlagsContractAddress<a id='EVM-FlagsContractAddress'></a>
```toml
FlagsContractAddress = '0xae4E781a6218A8031764928E88d457937A954fC3' # Example
```
**ADVANCED**

FlagsContractAddress can optionally point to a [Flags contract](../contracts/src/v0.8/Flags.sol). If set, the node will lookup that contract for each job that supports flags contracts (currently OCR and FM jobs are supported). If the job's contractAddress is set as hibernating in the FlagsContractAddress address, it overrides the standard update parameters (such as heartbeat/threshold).

### GasBumpPercent<a id='EVM-GasBumpPercent'></a>
```toml
GasBumpPercent = 10
```
The percentage by which to bump gas on a transaction that has exceeded `GasBumpThreshold`. The larger of `GasBumpPercent` and `GasBumpWei` is taken for gas bumps.

### GasBumpThreshold<a id='EVM-GasBumpThreshold'></a>
```toml
GasBumpThreshold = 15
```
Chainlink nodes can be configured to automatically bump gas on transactions that have been stuck waiting in the mempool for at least this many blocks. Set to 0 to disable gas bumping completely.

### GasBumpTxDepth<a id='EVM-GasBumpTxDepth'></a>
```toml
GasBumpTxDepth = 6
```
The number of transactions to gas bump starting from oldest. Set to 0 for no limit (i.e. bump all).

### GasBumpWei<a id='EVM-GasBumpWei'></a>
```toml
GasBumpWei = '100'
```
The minimum fixed amount of wei by which gas is bumped on each transaction attempt.

### GasEstimatorMode<a id='EVM-GasEstimatorMode'></a>
```toml
GasEstimatorMode = 'BlockHistory' # Default
```
Controls what type of gas estimator is used.

- `FixedPrice` uses static configured values for gas price (can be set via API call).
- `BlockHistory` dynamically adjusts default gas price based on heuristics from mined blocks.
- `L2Suggested`

Chainlink nodes decide what gas price to use using an `Estimator`. It ships with several simple and battle-hardened built-in estimators that should work well for almost all use-cases. Note that estimators will change their behaviour slightly depending on if you are in EIP-1559 mode or not.

You can also use your own estimator for gas price by selecting the `FixedPrice` estimator and using the exposed API to set the price.

An important point to note is that the Chainlink node does _not_ ship with built-in support for go-ethereum's `estimateGas` call. This is for several reasons, including security and reliability. We have found empirically that it is not generally safe to rely on the remote ETH node's idea of what gas price should be.

### GasFeeCapDefault<a id='EVM-GasFeeCapDefault'></a>
```toml
GasFeeCapDefault = '9223372036854775807'
```
If EIP1559 mode is enabled, and FixedPrice gas estimator is used, this env var controls the fixed initial fee cap.

### GasLimitDefault<a id='EVM-GasLimitDefault'></a>
```toml
GasLimitDefault = '12'
```
The default gas limit for outgoing transactions. This should not need to be changed in most cases.
Some job types, such as Keeper jobs, might set their own gas limit unrelated to this value.

### GasLimitMultiplier<a id='EVM-GasLimitMultiplier'></a>
```toml
GasLimitMultiplier = '1.234'
```
A factor by which a transaction's GasLimit is multiplied before transmission. So if the value is 1.1, and the GasLimit for a transaction is 10, 10% will be added before transmission.

This factor is always applied, so includes Optimism L2 transactions which uses a default gas limit of 1 and is also applied to EthGasLimitDefault.

### GasLimitTransfer<a id='EVM-GasLimitTransfer'></a>
```toml
GasLimitTransfer = '100'
```
The gas limit used for an ordinary ETH transfer.

### GasPriceDefault<a id='EVM-GasPriceDefault'></a>
```toml
GasPriceDefault = '9223372036854775807'
```
(Only applies to legacy transactions)

The default gas price to use when submitting transactions to the blockchain. Will be overridden by the built-in `BlockHistoryEstimator` if enabled, and might be increased if gas bumping is enabled.

Can be used with the `chainlink setgasprice` to be updated while the node is still running.

### GasTipCapDefault<a id='EVM-GasTipCapDefault'></a>
```toml
GasTipCapDefault = '2'
```
(Only applies to EIP-1559 transactions)

The default gas tip to use when submitting transactions to the blockchain. Will be overridden by the built-in `BlockHistoryEstimator` if enabled, and might be increased if gas bumping is enabled.

### GasTipCapMinimum<a id='EVM-GasTipCapMinimum'></a>
```toml
GasTipCapMinimum = '1'
```
Only applies to EIP-1559 transactions)

The minimum gas tip to use when submitting transactions to the blockchain.

### LinkContractAddress<a id='EVM-LinkContractAddress'></a>
```toml
LinkContractAddress = '0x538aAaB4ea120b2bC2fe5D296852D948F07D849e'
```


### LogBackfillBatchSize<a id='EVM-LogBackfillBatchSize'></a>
```toml
LogBackfillBatchSize = 17
```


### LogPollInterval<a id='EVM-LogPollInterval'></a>
```toml
LogPollInterval = '1m'
```


### MaxGasPriceWei<a id='EVM-MaxGasPriceWei'></a>
```toml
MaxGasPriceWei = '281474976710655'
```
Chainlink nodes will never pay more than this for a transaction.

### MaxInFlightTransactions<a id='EVM-MaxInFlightTransactions'></a>
```toml
MaxInFlightTransactions = 19
```
Controls how many transactions are allowed to be "in-flight" i.e. broadcast but unconfirmed at any one time. You can consider this a form of transaction throttling.

The default is set conservatively at 16 because this is a pessimistic minimum that both geth and parity will hold without evicting local transactions. If your node is falling behind and you need higher throughput, you can increase this setting, but you MUST make sure that your ETH node is configured properly otherwise you can get nonce gapped and your node will get stuck.

0 value disables the limit. Use with caution.

### MaxQueuedTransactions<a id='EVM-MaxQueuedTransactions'></a>
```toml
MaxQueuedTransactions = 99
```
The maximum number of unbroadcast transactions per key that are allowed to be enqueued before jobs will start failing and rejecting send of any further transactions. This represents a sanity limit and generally indicates a problem with your ETH node (transactions are not getting mined).

Do NOT blindly increase this value thinking it will fix things if you start hitting this limit because transactions are not getting mined, you will instead only make things worse.

In deployments with very high burst rates, or on chains with large re-orgs, you _may_ consider increasing this.

0 value disables any limit on queue size. Use with caution.

### MinGasPriceWei<a id='EVM-MinGasPriceWei'></a>
```toml
MinGasPriceWei = '13'
```
(Only applies to legacy transactions)

Chainlink nodes will never pay less than this for a transaction.

It is possible to force the Chainlink node to use a fixed gas price by setting a combination of these, e.g.

```toml
EIP1559DynamicFees = false
MaxGasPriceWei = 100
MinGasPriceWei = 100
GasPriceDefault = 100
GasBumpThreshold = 0
GasEstimatorMode = 'FixedPrice'
```

### MinIncomingConfirmations<a id='EVM-MinIncomingConfirmations'></a>
```toml
MinIncomingConfirmations = 13
```


### MinimumContractPayment<a id='EVM-MinimumContractPayment'></a>
```toml
MinimumContractPayment = '9223372036854775807'
```


### NonceAutoSync<a id='EVM-NonceAutoSync'></a>
```toml
NonceAutoSync = false # Default
```
Chainlink nodes will automatically try to sync its local nonce with the remote chain on startup and fast forward if necessary. This is almost always safe but can be disabled in exceptional cases by setting this value to false.

### OCRContractConfirmations<a id='EVM-OCRContractConfirmations'></a>
```toml
OCRContractConfirmations = 11
```


### OCRContractTransmitterTransmitTimeout<a id='EVM-OCRContractTransmitterTransmitTimeout'></a>
```toml
OCRContractTransmitterTransmitTimeout = '1m'
```


### OCRDatabaseTimeout<a id='EVM-OCRDatabaseTimeout'></a>
```toml
OCRDatabaseTimeout = '1s'
```


### OCRObservationGracePeriod<a id='EVM-OCRObservationGracePeriod'></a>
```toml
OCRObservationGracePeriod = '1s'
```


### OCRObservationTimeout<a id='EVM-OCRObservationTimeout'></a>
```toml
OCRObservationTimeout = '1m'
```


### OCR2ContractConfirmations<a id='EVM-OCR2ContractConfirmations'></a>
```toml
OCR2ContractConfirmations = 7
```


### OperatorFactoryAddress<a id='EVM-OperatorFactoryAddress'></a>
```toml
OperatorFactoryAddress = '0xa5B85635Be42F21f94F28034B7DA440EeFF0F418'
```


### RPCDefaultBatchSize<a id='EVM-RPCDefaultBatchSize'></a>
```toml
RPCDefaultBatchSize = 17
```


### TxReaperInterval<a id='EVM-TxReaperInterval'></a>
```toml
TxReaperInterval = '1m'
```


### TxReaperThreshold<a id='EVM-TxReaperThreshold'></a>
```toml
TxReaperThreshold = '1m'
```


### TxResendAfterThreshold<a id='EVM-TxResendAfterThreshold'></a>
```toml
TxResendAfterThreshold = '1h'
```


### UseForwarders<a id='EVM-UseForwarders'></a>
```toml
UseForwarders = false # Default
```
Enables or disables sending transactions through forwarder contracts.

## EVM.BlockHistoryEstimator<a id='EVM-BlockHistoryEstimator'></a>
```toml
[EVM.BlockHistoryEstimator]
BatchSize = 17
BlockDelay = 10
BlockHistorySize = 12
EIP1559FeeCapBufferBlocks = 13
TransactionPercentile = 15
```
These settings allow you to configure how your node calculates gas prices. In most cases, leaving these values at their defaults should give good results.

### BatchSize<a id='EVM-BlockHistoryEstimator-BatchSize'></a>
```toml
BatchSize = 17
```
Sets the maximum number of blocks to fetch in one batch in the block history estimator.
If the `BLOCK_HISTORY_ESTIMATOR_BATCH_SIZE` environment variable is set to 0, it defaults to ETH_RPC_DEFAULT_BATCH_SIZE.

### BlockDelay<a id='EVM-BlockHistoryEstimator-BlockDelay'></a>
```toml
BlockDelay = 10
```
Controls the number of blocks that the block history estimator trails behind head.
For example, if this is set to 3, and we receive block 10, block history estimator will fetch block 7.

CAUTION: You might be tempted to set this to 0 to use the latest possible
block, but it is possible to receive a head BEFORE that block is actually
available from the connected node via RPC, due to race conditions in the code of the remote ETH node. In this case you will get false
"zero" blocks that are missing transactions.

### BlockHistorySize<a id='EVM-BlockHistoryEstimator-BlockHistorySize'></a>
```toml
BlockHistorySize = 12
```
Controls the number of past blocks to keep in memory to use as a basis for calculating a percentile gas price.

### EIP1559FeeCapBufferBlocks<a id='EVM-BlockHistoryEstimator-EIP1559FeeCapBufferBlocks'></a>
```toml
EIP1559FeeCapBufferBlocks = 13
```
**ADVANCED**

If EIP1559 mode is enabled, this optional env var controls the buffer blocks to add to the current base fee when sending a transaction. By default, the gas bumping threshold + 1 block is used. It is not recommended to change this unless you know what you are doing.

### TransactionPercentile<a id='EVM-BlockHistoryEstimator-TransactionPercentile'></a>
```toml
TransactionPercentile = 15
```
Must be in range 0-100.

Only has an effect if gas updater is enabled. Specifies percentile gas price to choose. E.g. if the block history contains four transactions with gas prices `[100, 200, 300, 400]` then picking 25 for this number will give a value of 200. If the calculated gas price is higher than `ETH_GAS_PRICE_DEFAULT` then the higher price will be used as the base price for new transactions.

Think of this number as an indicator of how aggressive you want your node to price its transactions.

Setting this number higher will cause the Chainlink node to select higher gas prices.

Setting it lower will tend to set lower gas prices.

## EVM.HeadTracker<a id='EVM-HeadTracker'></a>
```toml
[EVM.HeadTracker]
BlockEmissionIdleWarningThreshold = '1m' # Example
HistoryDepth = 15
MaxBufferSize = 17
SamplingInterval = '1h'
```


### BlockEmissionIdleWarningThreshold<a id='EVM-HeadTracker-BlockEmissionIdleWarningThreshold'></a>
```toml
BlockEmissionIdleWarningThreshold = '1m' # Example
```
BlockEmissionIdleWarningThreshold will cause Chainlink to log warnings if this duration is exceeded without any new blocks being emitted.

### HistoryDepth<a id='EVM-HeadTracker-HistoryDepth'></a>
```toml
HistoryDepth = 15
```


### MaxBufferSize<a id='EVM-HeadTracker-MaxBufferSize'></a>
```toml
MaxBufferSize = 17
```


### SamplingInterval<a id='EVM-HeadTracker-SamplingInterval'></a>
```toml
SamplingInterval = '1h'
```


## EVM.KeySpecific<a id='EVM-KeySpecific'></a>
```toml
[[EVM.KeySpecific]]
Key = '0x2a3e23c6f242F5345320814aC8a1b4E58707D292'
MaxGasPriceWei = '79228162514264337593543950335'
```


### Key<a id='EVM-KeySpecific-Key'></a>
```toml
Key = '0x2a3e23c6f242F5345320814aC8a1b4E58707D292'
```


### MaxGasPriceWei<a id='EVM-KeySpecific-MaxGasPriceWei'></a>
```toml
MaxGasPriceWei = '79228162514264337593543950335'
```


## EVM.NodePool<a id='EVM-NodePool'></a>
```toml
[EVM.NodePool]
NoNewHeadsThreshold = '3m' # Default
PollFailureThreshold = 3 # Default
PollInterval = '10s' # Default
```


### NoNewHeadsThreshold<a id='EVM-NodePool-NoNewHeadsThreshold'></a>
```toml
NoNewHeadsThreshold = '3m' # Default
```
Controls how long to wait after receiving no new heads before marking the node as out-of-sync.

Set to zero to disable out-of-sync checking.

### PollFailureThreshold<a id='EVM-NodePool-PollFailureThreshold'></a>
```toml
PollFailureThreshold = 3 # Default
```
Indicates how many consecutive polls must fail in order to mark a node as unreachable.

Set to zero to disable poll checking.

### PollInterval<a id='EVM-NodePool-PollInterval'></a>
```toml
PollInterval = '10s' # Default
```
Controls how often to poll the node to check for liveness.

Set to zero to disable poll checking.

## EVM.Nodes<a id='EVM-Nodes'></a>
```toml
[[EVM.Nodes]]
Name = 'foo'
WSURL = 'wss://web.socket/test'
HTTPURL = 'https://foo.web'
SendOnly = true
```


### Name<a id='EVM-Nodes-Name'></a>
```toml
Name = 'foo'
```


### WSURL<a id='EVM-Nodes-WSURL'></a>
```toml
WSURL = 'wss://web.socket/test'
```


### HTTPURL<a id='EVM-Nodes-HTTPURL'></a>
```toml
HTTPURL = 'https://foo.web'
```


### SendOnly<a id='EVM-Nodes-SendOnly'></a>
```toml
SendOnly = true
```


## Solana<a id='Solana'></a>
```toml
[[Solana]]
ChainID = 'mainnet'
Enabled = false # Default
BalancePollPeriod = '5s' # Default
ConfirmPollPeriod = '500ms' # Default
OCR2CachePollPeriod = '1s' # Default
OCR2CacheTTL = '1m' # Default
TxTimeout = '1h' # Default
TxRetryTimeout = '10s' # Default
TxConfirmTimeout = '30s' # Default
SkipPreflight = true # Default
Commitment = 'confirmed' # Default
MaxRetries = 0 # Default
```


### ChainID<a id='Solana-ChainID'></a>
```toml
ChainID = 'mainnet'
```


### Enabled<a id='Solana-Enabled'></a>
```toml
Enabled = false # Default
```


### BalancePollPeriod<a id='Solana-BalancePollPeriod'></a>
```toml
BalancePollPeriod = '5s' # Default
```
BalancePollPeriod is the rate to poll for SOL balance and update Prometheus metrics.

### ConfirmPollPeriod<a id='Solana-ConfirmPollPeriod'></a>
```toml
ConfirmPollPeriod = '500ms' # Default
```
ConfirmPollPeriod is the rate to poll for signature confirmation.

### OCR2CachePollPeriod<a id='Solana-OCR2CachePollPeriod'></a>
```toml
OCR2CachePollPeriod = '1s' # Default
```
OCR2CachePollPeriod is the rate to poll for the OCR2 state cache.

### OCR2CacheTTL<a id='Solana-OCR2CacheTTL'></a>
```toml
OCR2CacheTTL = '1m' # Default
```
OCR2CacheTTl is the stale OCR2 cache deadline.

### TxTimeout<a id='Solana-TxTimeout'></a>
```toml
TxTimeout = '1h' # Default
```
TxTimeout is the timeout for sending txes to an RPC endpoint.

### TxRetryTimeout<a id='Solana-TxRetryTimeout'></a>
```toml
TxRetryTimeout = '10s' # Default
```
TxRetryTimeout is the duration for tx manager to attempt rebroadcasting to RPC, before giving up.

### TxConfirmTimeout<a id='Solana-TxConfirmTimeout'></a>
```toml
TxConfirmTimeout = '30s' # Default
```
TxConfirmTimeout is the duration to wait when confirming a tx signature, before discarding as unconfirmed.

### SkipPreflight<a id='Solana-SkipPreflight'></a>
```toml
SkipPreflight = true # Default
```
SkipPreflight enables or disables preflight checks when sending txs.

### Commitment<a id='Solana-Commitment'></a>
```toml
Commitment = 'confirmed' # Default
```
Confirmation level for solana state and transactions. ([documentation](https://docs.solana.com/developing/clients/jsonrpc-api#configuring-state-commitment))

### MaxRetries<a id='Solana-MaxRetries'></a>
```toml
MaxRetries = 0 # Default
```
When sending transactions, how many times the RPC node will automatically rebroadcast a tx.
The default is 0 for custom txm rebroadcasting method, set to -1 to use the RPC node's default retry strategy.

## Solana.Nodes<a id='Solana-Nodes'></a>
```toml
[[Solana.Nodes]]
Name = 'primary'
URL = 'http://solana.web'
```


### Name<a id='Solana-Nodes-Name'></a>
```toml
Name = 'primary'
```


### URL<a id='Solana-Nodes-URL'></a>
```toml
URL = 'http://solana.web'
```


## Terra<a id='Terra'></a>
```toml
[[Terra]]
ChainID = 'Bombay-12'
Enabled = true # Default
BlockRate = '1m'
BlocksUntilTxTimeout = 12
ConfirmPollPeriod = '1s'
FallbackGasPriceULuna = '0.001'
FCDURL = 'http://terra.com'
GasLimitMultiplier = '1.2'
MaxMsgsPerBatch = 17
OCR2CachePollPeriod = '1m'
OCR2CacheTTL = '1h'
TxMsgTimeout = '1s'
```


### ChainID<a id='Terra-ChainID'></a>
```toml
ChainID = 'Bombay-12'
```


### Enabled<a id='Terra-Enabled'></a>
```toml
Enabled = true # Default
```


### BlockRate<a id='Terra-BlockRate'></a>
```toml
BlockRate = '1m'
```


### BlocksUntilTxTimeout<a id='Terra-BlocksUntilTxTimeout'></a>
```toml
BlocksUntilTxTimeout = 12
```


### ConfirmPollPeriod<a id='Terra-ConfirmPollPeriod'></a>
```toml
ConfirmPollPeriod = '1s'
```


### FallbackGasPriceULuna<a id='Terra-FallbackGasPriceULuna'></a>
```toml
FallbackGasPriceULuna = '0.001'
```


### FCDURL<a id='Terra-FCDURL'></a>
```toml
FCDURL = 'http://terra.com'
```


### GasLimitMultiplier<a id='Terra-GasLimitMultiplier'></a>
```toml
GasLimitMultiplier = '1.2'
```


### MaxMsgsPerBatch<a id='Terra-MaxMsgsPerBatch'></a>
```toml
MaxMsgsPerBatch = 17
```


### OCR2CachePollPeriod<a id='Terra-OCR2CachePollPeriod'></a>
```toml
OCR2CachePollPeriod = '1m'
```


### OCR2CacheTTL<a id='Terra-OCR2CacheTTL'></a>
```toml
OCR2CacheTTL = '1h'
```


### TxMsgTimeout<a id='Terra-TxMsgTimeout'></a>
```toml
TxMsgTimeout = '1s'
```


## Terra.Nodes<a id='Terra-Nodes'></a>
```toml
[[Terra.Nodes]]
Name = 'primary'
TendermintURL = 'http://tender.mint'
```


### Name<a id='Terra-Nodes-Name'></a>
```toml
Name = 'primary'
```


### TendermintURL<a id='Terra-Nodes-TendermintURL'></a>
```toml
TendermintURL = 'http://tender.mint'
```


