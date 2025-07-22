package db

import "time"

// ResourceCacheUser designates the column to set in the resource_cache_users
// table.
type ResourceCacheUser interface {
	SQLMap() map[string]any
}

type forBuild struct {
	BuildID int
}

func ForBuild(id int) ResourceCacheUser {
	return forBuild{id}
}

func (user forBuild) SQLMap() map[string]any {
	return map[string]any{
		"build_id": user.BuildID,
	}
}

type forInMemoryBuild struct {
	BuildID    int
	CreateTime time.Time
}

func ForInMemoryBuild(id int, createTime time.Time) ResourceCacheUser {
	return forInMemoryBuild{id, createTime}
}

func (user forInMemoryBuild) SQLMap() map[string]any {
	return map[string]any{
		"in_memory_build_id":          user.BuildID,
		"in_memory_build_create_time": user.CreateTime,
	}
}
