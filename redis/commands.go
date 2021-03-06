package redis

var commands = []RedisHelp{
	{Command: "APPEND", Args: "key value", Desc: "Append a value to a key"},
	{Command: "AUTH", Args: "password", Desc: "Authenticate to the server"},
	{Command: "BGREWRITEAOF", Args: "-", Desc: "Asynchronously rewrite the append-only file"},
	{Command: "BGSAVE", Args: "-", Desc: "Asynchronously save the dataset to disk"},
	{Command: "BITCOUNT", Args: "key [start end]", Desc: "Count set bits in a string"},
	{Command: "BITFIELD", Args: "key [GET type offset] [SET type offset value] [INCRBY type offset increment] [OVERFLOW WRAP|SAT|FAIL]", Desc: "Perform arbitrary bitfield integer operations on strings"},
	{Command: "BITOP", Args: "operation destkey key [key ...]", Desc: "Perform bitwise operations between strings"},
	{Command: "BITPOS", Args: "key bit [start] [end]", Desc: "Find first bit set or clear in a string"},
	{Command: "BLPOP", Args: "key [key ...] timeout", Desc: "Remove and get the first element in a list, or block until one is available"},
	{Command: "BRPOP", Args: "key [key ...] timeout", Desc: "Remove and get the last element in a list, or block until one is available"},
	{Command: "BRPOPLPUSH", Args: "source destination timeout", Desc: "Pop a value from a list, push it to another list and return it; or block until one is available"},
	{Command: "BZPOPMAX", Args: "key [key ...] timeout", Desc: "Remove and return the member with the highest score from one or more sorted sets, or block until one is available"},
	{Command: "BZPOPMIN", Args: "key [key ...] timeout", Desc: "Remove and return the member with the lowest score from one or more sorted sets, or block until one is available"},
	{Command: "CLIENT GETNAME", Args: "-", Desc: "Get the current connection name"},
	{Command: "CLIENT ID", Args: "-", Desc: "Returns the client ID for the current connection"},
	{Command: "CLIENT KILL", Args: "[ip:port] [ID client-id] [TYPE normal|master|slave|pubsub] [ADDR ip:port] [SKIPME yes/no]", Desc: "Kill the connection of a client"},
	{Command: "CLIENT LIST", Args: "-", Desc: "Get the list of client connections"},
	{Command: "CLIENT PAUSE", Args: "timeout", Desc: "Stop processing commands from clients for some time"},
	{Command: "CLIENT REPLY", Args: "ON|OFF|SKIP", Desc: "Instruct the server whether to reply to commands"},
	{Command: "CLIENT SETNAME", Args: "connection-name", Desc: "Set the current connection name"},
	{Command: "CLIENT UNBLOCK", Args: "client-id [TIMEOUT|ERROR]", Desc: "Unblock a client blocked in a blocking command from a different connection"},
	{Command: "CLUSTER ADDSLOTS", Args: "slot [slot ...]", Desc: "Assign new hash slots to receiving node"},
	{Command: "CLUSTER COUNT-FAILURE-REPORTS", Args: "node-id", Desc: "Return the number of failure reports active for a given node"},
	{Command: "CLUSTER COUNTKEYSINSLOT", Args: "slot", Desc: "Return the number of local keys in the specified hash slot"},
	{Command: "CLUSTER DELSLOTS", Args: "slot [slot ...]", Desc: "Set hash slots as unbound in receiving node"},
	{Command: "CLUSTER FAILOVER", Args: "[FORCE|TAKEOVER]", Desc: "Forces a replica to perform a manual failover of its master."},
	{Command: "CLUSTER FORGET", Args: "node-id", Desc: "Remove a node from the nodes table"},
	{Command: "CLUSTER GETKEYSINSLOT", Args: "slot count", Desc: "Return local key names in the specified hash slot"},
	{Command: "CLUSTER INFO", Args: "-", Desc: "Provides info about Redis Cluster node state"},
	{Command: "CLUSTER KEYSLOT", Args: "key", Desc: "Returns the hash slot of the specified key"},
	{Command: "CLUSTER MEET", Args: "ip port", Desc: "Force a node cluster to handshake with another node"},
	{Command: "CLUSTER NODES", Args: "-", Desc: "Get Cluster config for the node"},
	{Command: "CLUSTER REPLICAS", Args: "node-id", Desc: "List replica nodes of the specified master node"},
	{Command: "CLUSTER REPLICATE", Args: "node-id", Desc: "Reconfigure a node as a replica of the specified master node"},
	{Command: "CLUSTER RESET", Args: "[HARD|SOFT]", Desc: "Reset a Redis Cluster node"},
	{Command: "CLUSTER SAVECONFIG", Args: "-", Desc: "Forces the node to save cluster state on disk"},
	{Command: "CLUSTER SET-CONFIG-EPOCH", Args: "config-epoch", Desc: "Set the configuration epoch in a new node"},
	{Command: "CLUSTER SETSLOT", Args: "slot IMPORTING|MIGRATING|STABLE|NODE [node-id]", Desc: "Bind a hash slot to a specific node"},
	{Command: "CLUSTER SLAVES", Args: "node-id", Desc: "List replica nodes of the specified master node"},
	{Command: "CLUSTER SLOTS", Args: "-", Desc: "Get array of Cluster slot to node mappings"},
	{Command: "COMMAND", Args: "-", Desc: "Get array of Redis command details"},
	{Command: "COMMAND COUNT", Args: "-", Desc: "Get total number of Redis commands"},
	{Command: "COMMAND GETKEYS", Args: "-", Desc: "Extract keys given a full Redis command"},
	{Command: "COMMAND INFO", Args: "command-name [command-name ...]", Desc: "Get array of specific Redis command details"},
	{Command: "CONFIG GET", Args: "parameter", Desc: "Get the value of a configuration parameter"},
	{Command: "CONFIG RESETSTAT", Args: "-", Desc: "Reset the stats returned by INFO"},
	{Command: "CONFIG REWRITE", Args: "-", Desc: "Rewrite the configuration file with the in memory configuration"},
	{Command: "CONFIG SET", Args: "parameter value", Desc: "Set a configuration parameter to the given value"},
	{Command: "DBSIZE", Args: "-", Desc: "Return the number of keys in the selected database"},
	{Command: "DEBUG OBJECT", Args: "key", Desc: "Get debugging information about a key"},
	{Command: "DEBUG SEGFAULT", Args: "-", Desc: "Make the server crash"},
	{Command: "DECR", Args: "key", Desc: "Decrement the integer value of a key by one"},
	{Command: "DECRBY", Args: "key decrement", Desc: "Decrement the integer value of a key by the given number"},
	{Command: "DEL", Args: "key [key ...]", Desc: "Delete a key"},
	{Command: "DISCARD", Args: "-", Desc: "Discard all commands issued after MULTI"},
	{Command: "DUMP", Args: "key", Desc: "Return a serialized version of the value stored at the specified key."},
	{Command: "ECHO", Args: "message", Desc: "Echo the given string"},
	{Command: "EVAL", Args: "script numkeys key [key ...] arg [arg ...]", Desc: "Execute a Lua script server side"},
	{Command: "EVALSHA", Args: "sha1 numkeys key [key ...] arg [arg ...]", Desc: "Execute a Lua script server side"},
	{Command: "EXEC", Args: "-", Desc: "Execute all commands issued after MULTI"},
	{Command: "EXISTS", Args: "key [key ...]", Desc: "Determine if a key exists"},
	{Command: "EXPIRE", Args: "key seconds", Desc: "Set a key's time to live in seconds"},
	{Command: "EXPIREAT", Args: "key timestamp", Desc: "Set the expiration for a key as a UNIX timestamp"},
	{Command: "FLUSHALL", Args: "[ASYNC]", Desc: "Remove all keys from all databases"},
	{Command: "FLUSHDB", Args: "[ASYNC]", Desc: "Remove all keys from the current database"},
	{Command: "GEOADD", Args: "key longitude latitude member [longitude latitude member ...]", Desc: "Add one or more geospatial items in the geospatial index represented using a sorted set"},
	{Command: "GEODIST", Args: "key member1 member2 [unit]", Desc: "Returns the distance between two members of a geospatial index"},
	{Command: "GEOHASH", Args: "key member [member ...]", Desc: "Returns members of a geospatial index as standard geohash strings"},
	{Command: "GEOPOS", Args: "key member [member ...]", Desc: "Returns longitude and latitude of members of a geospatial index"},
	{Command: "GEORADIUS", Args: "key longitude latitude radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count] [ASC|DESC] [STORE key] [STOREDIST key]", Desc: "Query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a point"},
	{Command: "GEORADIUSBYMEMBER", Args: "key member radius m|km|ft|mi [WITHCOORD] [WITHDIST] [WITHHASH] [COUNT count] [ASC|DESC] [STORE key] [STOREDIST key]", Desc: "Query a sorted set representing a geospatial index to fetch members matching a given maximum distance from a member"},
	{Command: "GET", Args: "key", Desc: "Get the value of a key"},
	{Command: "GETBIT", Args: "key offset", Desc: "Returns the bit value at offset in the string value stored at key"},
	{Command: "GETRANGE", Args: "key start end", Desc: "Get a substring of the string stored at a key"},
	{Command: "GETSET", Args: "key value", Desc: "Set the string value of a key and return its old value"},
	{Command: "HDEL", Args: "key field [field ...]", Desc: "Delete one or more hash fields"},
	{Command: "HEXISTS", Args: "key field", Desc: "Determine if a hash field exists"},
	{Command: "HGET", Args: "key field", Desc: "Get the value of a hash field"},
	{Command: "HGETALL", Args: "key", Desc: "Get all the fields and values in a hash"},
	{Command: "HINCRBY", Args: "key field increment", Desc: "Increment the integer value of a hash field by the given number"},
	{Command: "HINCRBYFLOAT", Args: "key field increment", Desc: "Increment the float value of a hash field by the given amount"},
	{Command: "HKEYS", Args: "key", Desc: "Get all the fields in a hash"},
	{Command: "HLEN", Args: "key", Desc: "Get the number of fields in a hash"},
	{Command: "HMGET", Args: "key field [field ...]", Desc: "Get the values of all the given hash fields"},
	{Command: "HMSET", Args: "key field value [field value ...]", Desc: "Set multiple hash fields to multiple values"},
	{Command: "HSCAN", Args: "key cursor [MATCH pattern] [COUNT count]", Desc: "Incrementally iterate hash fields and associated values"},
	{Command: "HSET", Args: "key field value", Desc: "Set the string value of a hash field"},
	{Command: "HSETNX", Args: "key field value", Desc: "Set the value of a hash field, only if the field does not exist"},
	{Command: "HSTRLEN", Args: "key field", Desc: "Get the length of the value of a hash field"},
	{Command: "HVALS", Args: "key", Desc: "Get all the values in a hash"},
	{Command: "INCR", Args: "key", Desc: "Increment the integer value of a key by one"},
	{Command: "INCRBY", Args: "key increment", Desc: "Increment the integer value of a key by the given amount"},
	{Command: "INCRBYFLOAT", Args: "key increment", Desc: "Increment the float value of a key by the given amount"},
	{Command: "INFO", Args: "[section]", Desc: "Get information and statistics about the server"},
	{Command: "KEYS", Args: "pattern", Desc: "Find all keys matching the given pattern"},
	{Command: "LASTSAVE", Args: "-", Desc: "Get the UNIX time stamp of the last successful save to disk"},
	{Command: "LINDEX", Args: "key index", Desc: "Get an element from a list by its index"},
	{Command: "LINSERT", Args: "key BEFORE|AFTER pivot value", Desc: "Insert an element before or after another element in a list"},
	{Command: "LLEN", Args: "key", Desc: "Get the length of a list"},
	{Command: "LPOP", Args: "key", Desc: "Remove and get the first element in a list"},
	{Command: "LPUSH", Args: "key value [value ...]", Desc: "Prepend one or multiple values to a list"},
	{Command: "LPUSHX", Args: "key value", Desc: "Prepend a value to a list, only if the list exists"},
	{Command: "LRANGE", Args: "key start stop", Desc: "Get a range of elements from a list"},
	{Command: "LREM", Args: "key count value", Desc: "Remove elements from a list"},
	{Command: "LSET", Args: "key index value", Desc: "Set the value of an element in a list by its index"},
	{Command: "LTRIM", Args: "key start stop", Desc: "Trim a list to the specified range"},
	{Command: "MEMORY DOCTOR", Args: "-", Desc: "Outputs memory problems report"},
	{Command: "MEMORY HELP", Args: "-", Desc: "Show helpful text about the different subcommands"},
	{Command: "MEMORY MALLOC-STATS", Args: "-", Desc: "Show allocator internal stats"},
	{Command: "MEMORY PURGE", Args: "-", Desc: "Ask the allocator to release memory"},
	{Command: "MEMORY STATS", Args: "-", Desc: "Show memory usage details"},
	{Command: "MEMORY USAGE", Args: "key [SAMPLES count]", Desc: "Estimate the memory usage of a key"},
	{Command: "MGET", Args: "key [key ...]", Desc: "Get the values of all the given keys"},
	{Command: "MIGRATE", Args: "host port key | destination-db timeout [COPY] [REPLACE] [KEYS key]", Desc: "Atomically transfer a key from a Redis instance to another one."},
	{Command: "MONITOR", Args: "-", Desc: "Listen for all requests received by the server in real time"},
	{Command: "MOVE", Args: "key db", Desc: "Move a key to another database"},
	{Command: "MSET", Args: "key value [key value ...]", Desc: "Set multiple keys to multiple values"},
	{Command: "MSETNX", Args: "key value [key value ...]", Desc: "Set multiple keys to multiple values, only if none of the keys exist"},
	{Command: "MULTI", Args: "-", Desc: "Mark the start of a transaction block"},
	{Command: "OBJECT", Args: "subcommand [arguments [arguments ...]]", Desc: "Inspect the internals of Redis objects"},
	{Command: "PERSIST", Args: "key", Desc: "Remove the expiration from a key"},
	{Command: "PEXPIRE", Args: "key milliseconds", Desc: "Set a key's time to live in milliseconds"},
	{Command: "PEXPIREAT", Args: "key milliseconds-timestamp", Desc: "Set the expiration for a key as a UNIX timestamp specified in milliseconds"},
	{Command: "PFADD", Args: "key element [element ...]", Desc: "Adds the specified elements to the specified HyperLogLog."},
	{Command: "PFCOUNT", Args: "key [key ...]", Desc: "Return the approximated cardinality of the set(s) observed by the HyperLogLog at key(s)."},
	{Command: "PFMERGE", Args: "destkey sourcekey [sourcekey ...]", Desc: "Merge N different HyperLogLogs into a single one."},
	{Command: "PING", Args: "[message]", Desc: "Ping the server"},
	{Command: "PSETEX", Args: "key milliseconds value", Desc: "Set the value and expiration in milliseconds of a key"},
	{Command: "PSUBSCRIBE", Args: "pattern [pattern ...]", Desc: "Listen for messages published to channels matching the given patterns"},
	{Command: "PTTL", Args: "key", Desc: "Get the time to live for a key in milliseconds"},
	{Command: "PUBLISH", Args: "channel message", Desc: "Post a message to a channel"},
	{Command: "PUBSUB", Args: "subcommand [argument [argument ...]]", Desc: "Inspect the state of the Pub/Sub subsystem"},
	{Command: "PUNSUBSCRIBE", Args: "[pattern [pattern ...]]", Desc: "Stop listening for messages posted to channels matching the given patterns"},
	{Command: "QUIT", Args: "-", Desc: "Close the connection"},
	{Command: "RANDOMKEY", Args: "-", Desc: "Return a random key from the keyspace"},
	{Command: "READONLY", Args: "-", Desc: "Enables read queries for a connection to a cluster replica node"},
	{Command: "READWRITE", Args: "-", Desc: "Disables read queries for a connection to a cluster replica node"},
	{Command: "RENAME", Args: "key newkey", Desc: "Rename a key"},
	{Command: "RENAMENX", Args: "key newkey", Desc: "Rename a key, only if the new key does not exist"},
	{Command: "REPLICAOF", Args: "host port", Desc: "Make the server a replica of another instance, or promote it as master."},
	{Command: "RESTORE", Args: "key ttl serialized-value [REPLACE]", Desc: "Create a key using the provided serialized value, previously obtained using DUMP."},
	{Command: "ROLE", Args: "-", Desc: "Return the role of the instance in the context of replication"},
	{Command: "RPOP", Args: "key", Desc: "Remove and get the last element in a list"},
	{Command: "RPOPLPUSH", Args: "source destination", Desc: "Remove the last element in a list, prepend it to another list and return it"},
	{Command: "RPUSH", Args: "key value [value ...]", Desc: "Append one or multiple values to a list"},
	{Command: "RPUSHX", Args: "key value", Desc: "Append a value to a list, only if the list exists"},
	{Command: "SADD", Args: "key member [member ...]", Desc: "Add one or more members to a set"},
	{Command: "SAVE", Args: "-", Desc: "Synchronously save the dataset to disk"},
	{Command: "SCAN", Args: "cursor [MATCH pattern] [COUNT count]", Desc: "Incrementally iterate the keys space"},
	{Command: "SCARD", Args: "key", Desc: "Get the number of members in a set"},
	{Command: "SCRIPT DEBUG", Args: "YES|SYNC|NO", Desc: "Set the debug mode for executed scripts."},
	{Command: "SCRIPT EXISTS", Args: "sha1 [sha1 ...]", Desc: "Check existence of scripts in the script cache."},
	{Command: "SCRIPT FLUSH", Args: "-", Desc: "Remove all the scripts from the script cache."},
	{Command: "SCRIPT KILL", Args: "-", Desc: "Kill the script currently in execution."},
	{Command: "SCRIPT LOAD", Args: "script", Desc: "Load the specified Lua script into the script cache."},
	{Command: "SDIFF", Args: "key [key ...]", Desc: "Subtract multiple sets"},
	{Command: "SDIFFSTORE", Args: "destination key [key ...]", Desc: "Subtract multiple sets and store the resulting set in a key"},
	{Command: "SELECT", Args: "index", Desc: "Change the selected database for the current connection"},
	{Command: "SET", Args: "key value [expiration EX seconds|PX milliseconds] [NX|XX]", Desc: "Set the string value of a key"},
	{Command: "SETBIT", Args: "key offset value", Desc: "Sets or clears the bit at offset in the string value stored at key"},
	{Command: "SETEX", Args: "key seconds value", Desc: "Set the value and expiration of a key"},
	{Command: "SETNX", Args: "key value", Desc: "Set the value of a key, only if the key does not exist"},
	{Command: "SETRANGE", Args: "key offset value", Desc: "Overwrite part of a string at key starting at the specified offset"},
	{Command: "SHUTDOWN", Args: "[NOSAVE|SAVE]", Desc: "Synchronously save the dataset to disk and then shut down the server"},
	{Command: "SINTER", Args: "key [key ...]", Desc: "Intersect multiple sets"},
	{Command: "SINTERSTORE", Args: "destination key [key ...]", Desc: "Intersect multiple sets and store the resulting set in a key"},
	{Command: "SISMEMBER", Args: "key member", Desc: "Determine if a given value is a member of a set"},
	{Command: "SLAVEOF", Args: "host port", Desc: "Make the server a replica of another instance, or promote it as master. Deprecated starting with Redis 5. Use REPLICAOF instead."},
	{Command: "SLOWLOG", Args: "subcommand [argument]", Desc: "Manages the Redis slow queries log"},
	{Command: "SMEMBERS", Args: "key", Desc: "Get all the members in a set"},
	{Command: "SMOVE", Args: "source destination member", Desc: "Move a member from one set to another"},
	{Command: "SORT", Args: "key [BY pattern] [LIMIT offset count] [GET pattern [GET pattern ...]] [ASC|DESC] [ALPHA] [STORE destination]", Desc: "Sort the elements in a list, set or sorted set"},
	{Command: "SPOP", Args: "key [count]", Desc: "Remove and return one or multiple random members from a set"},
	{Command: "SRANDMEMBER", Args: "key [count]", Desc: "Get one or multiple random members from a set"},
	{Command: "SREM", Args: "key member [member ...]", Desc: "Remove one or more members from a set"},
	{Command: "SSCAN", Args: "key cursor [MATCH pattern] [COUNT count]", Desc: "Incrementally iterate Set elements"},
	{Command: "STRLEN", Args: "key", Desc: "Get the length of the value stored in a key"},
	{Command: "SUBSCRIBE", Args: "channel [channel ...]", Desc: "Listen for messages published to the given channels"},
	{Command: "SUNION", Args: "key [key ...]", Desc: "Add multiple sets"},
	{Command: "SUNIONSTORE", Args: "destination key [key ...]", Desc: "Add multiple sets and store the resulting set in a key"},
	{Command: "SWAPDB", Args: "index index", Desc: "Swaps two Redis databases"},
	{Command: "SYNC", Args: "-", Desc: "Internal command used for replication"},
	{Command: "TIME", Args: "-", Desc: "Return the current server time"},
	{Command: "TOUCH", Args: "key [key ...]", Desc: "Alters the last access time of a key(s). Returns the number of existing keys specified."},
	{Command: "TTL", Args: "key", Desc: "Get the time to live for a key"},
	{Command: "TYPE", Args: "key", Desc: "Determine the type stored at key"},
	{Command: "UNLINK", Args: "key [key ...]", Desc: "Delete a key asynchronously in another thread. Otherwise it is just as DEL, but non blocking."},
	{Command: "UNSUBSCRIBE", Args: "[channel [channel ...]]", Desc: "Stop listening for messages posted to the given channels"},
	{Command: "UNWATCH", Args: "-", Desc: "Forget about all watched keys"},
	{Command: "WAIT", Args: "numreplicas timeout", Desc: "Wait for the synchronous replication of all the write commands sent in the context of the current connection"},
	{Command: "WATCH", Args: "key [key ...]", Desc: "Watch the given keys to determine execution of the MULTI/EXEC block"},
	{Command: "XACK", Args: "key group ID [ID ...]", Desc: "Marks a pending message as correctly processed, effectively removing it from the pending entries list of the consumer group. Return value of the command is the number of messages successfully acknowledged, that is, the IDs we were actually able to resolve in the PEL."},
	{Command: "XADD", Args: "key ID field string [field string ...]", Desc: "Appends a new entry to a stream"},
	{Command: "XCLAIM", Args: "key group consumer min-idle-time ID [ID ...] [IDLE ms] [TIME ms-unix-time] [RETRYCOUNT count] [force] [justid]", Desc: "Changes (or acquires) ownership of a message in a consumer group, as if the message was delivered to the specified consumer."},
	{Command: "XDEL", Args: "key ID [ID ...]", Desc: "Removes the specified entries from the stream. Returns the number of items actually deleted, that may be different from the number of IDs passed in case certain IDs do not exist."},
	{Command: "XGROUP", Args: "[CREATE key groupname id-or-$] [SETID key id-or-$] [DESTROY key groupname] [DELCONSUMER key groupname consumername]", Desc: "Create, destroy, and manage consumer groups."},
	{Command: "XINFO", Args: "[CONSUMERS key groupname] [GROUPS key] [STREAM key] [HELP]", Desc: "Get information on streams and consumer groups"},
	{Command: "XLEN", Args: "key", Desc: "Return the number of entires in a stream"},
	{Command: "XPENDING", Args: "key group [start end count] [consumer]", Desc: "Return information and entries from a stream consumer group pending entries list, that are messages fetched but never acknowledged."},
	{Command: "XRANGE", Args: "key start end [COUNT count]", Desc: "Return a range of elements in a stream, with IDs matching the specified IDs interval"},
	{Command: "XREAD", Args: "[COUNT count] [BLOCK milliseconds] STREAMS key [key ...] ID [ID ...]", Desc: "Return never seen elements in multiple streams, with IDs greater than the ones reported by the caller for each stream. Can block."},
	{Command: "XREADGROUP", Args: "GROUP group consumer [COUNT count] [BLOCK milliseconds] STREAMS key [key ...] ID [ID ...]", Desc: "Return new entries from a stream using a consumer group, or access the history of the pending entries for a given consumer. Can block."},
	{Command: "XREVRANGE", Args: "key end start [COUNT count]", Desc: "Return a range of elements in a stream, with IDs matching the specified IDs interval, in reverse order (from greater to smaller IDs) compared to XRANGE"},
	{Command: "XTRIM", Args: "key MAXLEN [~] count", Desc: "Trims the stream to (approximately if '~' is passed) a certain size"},
	{Command: "ZADD", Args: "key [NX|XX] [CH] [INCR] score member [score member ...]", Desc: "Add one or more members to a sorted set, or update its score if it already exists"},
	{Command: "ZCARD", Args: "key", Desc: "Get the number of members in a sorted set"},
	{Command: "ZCOUNT", Args: "key min max", Desc: "Count the members in a sorted set with scores within the given values"},
	{Command: "ZINCRBY", Args: "key increment member", Desc: "Increment the score of a member in a sorted set"},
	{Command: "ZINTERSTORE", Args: "destination numkeys key [key ...] [WEIGHTS weight] [AGGREGATE SUM|MIN|MAX]", Desc: "Intersect multiple sorted sets and store the resulting sorted set in a new key"},
	{Command: "ZLEXCOUNT", Args: "key min max", Desc: "Count the number of members in a sorted set between a given lexicographical range"},
	{Command: "ZPOPMAX", Args: "key [count]", Desc: "Remove and return members with the highest scores in a sorted set"},
	{Command: "ZPOPMIN", Args: "key [count]", Desc: "Remove and return members with the lowest scores in a sorted set"},
	{Command: "ZRANGE", Args: "key start stop [WITHSCORES]", Desc: "Return a range of members in a sorted set, by index"},
	{Command: "ZRANGEBYLEX", Args: "key min max [LIMIT offset count]", Desc: "Return a range of members in a sorted set, by lexicographical range"},
	{Command: "ZRANGEBYSCORE", Args: "key min max [WITHSCORES] [LIMIT offset count]", Desc: "Return a range of members in a sorted set, by score"},
	{Command: "ZRANK", Args: "key member", Desc: "Determine the index of a member in a sorted set"},
	{Command: "ZREM", Args: "key member [member ...]", Desc: "Remove one or more members from a sorted set"},
	{Command: "ZREMRANGEBYLEX", Args: "key min max", Desc: "Remove all members in a sorted set between the given lexicographical range"},
	{Command: "ZREMRANGEBYRANK", Args: "key start stop", Desc: "Remove all members in a sorted set within the given indexes"},
	{Command: "ZREMRANGEBYSCORE", Args: "key min max", Desc: "Remove all members in a sorted set within the given scores"},
	{Command: "ZREVRANGE", Args: "key start stop [WITHSCORES]", Desc: "Return a range of members in a sorted set, by index, with scores ordered from high to low"},
	{Command: "ZREVRANGEBYLEX", Args: "key max min [LIMIT offset count]", Desc: "Return a range of members in a sorted set, by lexicographical range, ordered from higher to lower strings."},
	{Command: "ZREVRANGEBYSCORE", Args: "key max min [WITHSCORES] [LIMIT offset count]", Desc: "Return a range of members in a sorted set, by score, with scores ordered from high to low"},
	{Command: "ZREVRANK", Args: "key member", Desc: "Determine the index of a member in a sorted set, with scores ordered from high to low"},
	{Command: "ZSCAN", Args: "key cursor [MATCH pattern] [COUNT count]", Desc: "Incrementally iterate sorted sets elements and associated scores"},
	{Command: "ZSCORE", Args: "key member", Desc: "Get the score associated with the given member in a sorted set"},
	{Command: "ZUNIONSTORE", Args: "destination numkeys key [key ...] [WEIGHTS weight] [AGGREGATE SUM|MIN|MAX]", Desc: "Add multiple sorted sets and store the resulting sorted set in a new key"},
}
