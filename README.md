bloodsport
==========

Bloodsport is a persistence engine inspired by redis written in Go.

Why?

Redis is a great concept but lacks practicality with large data sets.  It's primarily used by many organizations as a cache.

Redis cluster has been "coming soon" since 2011.

We're aiming to address the shortcomings of Redis, in particular:

1. Data size limits.  Redis doesn't allow you to go over your available memory.  This is pretty terrible.

2. Concurrency.  Redis consists of a single thread in order to serialize operations.  This means that the answer to concurrency was effectively to stick it's head in the ground and pretend it's not necessary.

3. Clustering.  We aim to build a clustered redis database server right out of the gate.  See the Kickboxer project, which will be handling the clustering aspects.  https://github.com/bdeggleston/kickboxer.  We will not be held back by the design limits of redis.



