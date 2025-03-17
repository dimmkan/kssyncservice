package sync

type SyncShedullerDeps struct {
	SyncRepository *SyncRepository
}

type SyncSheduller struct {
	SyncRepository *SyncRepository
}

func NewSyncSheduler(deps SyncShedullerDeps) *SyncSheduller {
	return &SyncSheduller{
		SyncRepository: deps.SyncRepository,
	}
}