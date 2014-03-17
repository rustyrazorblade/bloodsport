package commitlog


/*
It is recommended that the commit log be stored on a separate disk from the rest of the data files
this is because, like many dbs, the commitlog is append only.

Data should be writted to the commit log before being saved in memory, as it can be used to replay
 */

type CommitLog struct {

}
