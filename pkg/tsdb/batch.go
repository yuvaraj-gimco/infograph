package tsdb

import "sync"

type Batch struct {
	DataSourceId int64
	Queries      QuerySlice
	Depends      []*sync.WaitGroup
	Done         sync.WaitGroup
}

type BatchSlice []*Batch

func newBatch(dsId int64, queries QuerySlice) *Batch {
	return &Batch{
		DataSourceId: dsId,
		Queries:      queries,
	}
}

func (bg *Batch) process() error {
	for _, dep := range bg.Depends {
		dep.Wait()
	}
	return nil
}

func (bg *Batch) addQuery(query *Query) {
	bg.Queries = append(bg.Queries, query)
}

func getBatches(req *Request) (BatchSlice, error) {
	batches := make(BatchSlice, 0)

	for _, query := range req.Queries {
		if foundBatch := findMatchingBatchGroup(query, batches); foundBatch != nil {
			foundBatch.addQuery(query)
		} else {
			newBatch := newBatch(query.DataSource.Id, QuerySlice{query})
			batches = append(batches, newBatch)

			if query.DataSource.Meta {
				refIds := []string{"A", "B"}

				for _, refId := range refIds {
					for _, batch := range batches {
						for _, batchQuery := range batch.Queries {
							if batchQuery.RefId == refId {
								newBatch.Depends = append(newBatch.Depends, &batch.Done)
							}
						}
					}
				}
			}
		}
	}

	return batches, nil
}

func findMatchingBatchGroup(query *Query, batches BatchSlice) *Batch {
	for _, batch := range batches {
		if batch.DataSourceId == query.DataSource.Id {
			return batch
		}
	}
	return nil
}
