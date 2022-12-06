package service

import (
	"context"
	"database/sql"
	"time"

	api "github.com/anmi/go-ttsa/api"
	ttssqlc "github.com/anmi/go-ttsa/db"
	"github.com/anmi/go-ttsa/service/mapping"
	"github.com/anmi/go-ttsa/service/tsession.go"
	helpers "github.com/anmi/go-ttsa/service/utils"
)

func (s *TTApiService) CreateTask(ctx context.Context, req api.OptNewTaskForm) (api.CreateTaskRes, error) {
	session, session_err := helpers.GetSession(ctx, *s.Queries)
	if session_err != nil {
		return ServiceErrors.Unauthorized(), nil
	}

	if !session.RootTaskID.Valid {
		root_task_id, err := helpers.CreateRootTask(ctx, s.Queries, session.UserID)

		if err != nil {
			return &api.Task{}, err
		}

		session.RootTaskID = sql.NullInt64{Valid: true, Int64: root_task_id}
	}

	task, err := s.Queries.CreateTask(ctx, ttssqlc.CreateTaskParams{
		Title:       req.Value.Title,
		Description: req.Value.Description,
		Result:      "",
		CreatedAt:   time.Now(),
		CreatedBy:   session.UserID,
	})

	if err != nil {
		return &api.Task{}, err
	}

	parent_id := session.RootTaskID.Int64
	if req.Value.ParentId.IsSet() {
		parent_id = int64(req.Value.ParentId.Value)
	}

	_, err_task_dep := s.Queries.CreateTaskDependency(ctx, ttssqlc.CreateTaskDependencyParams{
		TaskID:      parent_id,
		DependsOnID: task.ID,
		CreatedAt:   time.Now(),
	})

	if err_task_dep != nil {
		return &api.Task{}, err_task_dep
	}

	return &api.Task{
		ID:          int(task.ID),
		Title:       task.Title,
		Description: task.Description,
		Result:      task.Result,
		CreatedAt:   task.CreatedAt.String(),
		DoneAt:      api.OptString{Set: false},
	}, nil
}

func (s *TTApiService) GetTask(ctx context.Context, params api.GetTaskParams) (api.GetTaskRes, error) {
	session, _ := tsession.GetTSession(ctx, s.Queries)
	task, task_err := session.GetTask(int64(params.ID))
	if task_err == tsession.SessionForbiddenError {
		return ServiceErrors.Forbidden(), nil
	}
	if task_err == tsession.SessionUnauthorizedError {
		return ServiceErrors.Unauthorized(), nil
	}
	if task_err != nil {
		return &api.Task{}, task_err
	}

	rtask := mapping.SqlTaskToApi(task)
	return &rtask, nil
}

func (s *TTApiService) GetTaskDependencies(ctx context.Context, params api.GetTaskDependenciesParams) (api.GetTaskDependenciesRes, error) {
	session, _ := tsession.GetTSession(ctx, s.Queries)
	_, task_err := session.GetTask(int64(params.ID))
	if task_err == tsession.SessionForbiddenError {
		return ServiceErrors.Forbidden(), nil
	}
	if task_err == tsession.SessionUnauthorizedError {
		return ServiceErrors.Unauthorized(), nil
	}
	if task_err != nil {
		return &api.TaskDependencies{}, task_err
	}

	tasks, err_tasks := s.Queries.GetTaskDependencies(ctx, int64(params.ID))

	if err_tasks != nil {
		return &api.TaskDependencies{}, err_tasks
	}

	apiTasks := []api.Task{}

	for _, task := range tasks {
		apiTasks = append(apiTasks, mapping.SqlTaskToApi(task))
	}

	return &api.TaskDependencies{
		Dependencies: apiTasks,
	}, nil
}

func (s *TTApiService) UpdateTask(ctx context.Context, req api.OptTaskUpdateForm, params api.UpdateTaskParams) (api.UpdateTaskRes, error) {
	session, _ := tsession.GetTSession(ctx, s.Queries)
	task, err := session.GetTask(int64(params.ID))
	if err == tsession.SessionForbiddenError {
		return ServiceErrors.Forbidden(), nil
	}
	if err == tsession.SessionUnauthorizedError {
		return ServiceErrors.Unauthorized(), nil
	}
	if err != nil {
		return &api.Task{}, err
	}

	if req.Set {
		form := req.Value

		if form.Title.Set {
			task, err = s.Queries.UpdateTaskTitle(ctx, ttssqlc.UpdateTaskTitleParams{
				Title: form.Title.Value,
				ID:    int64(params.ID),
			})

			if err != nil {
				return &api.Task{}, err
			}
		}

		if form.Description.Set {
			task, err = s.Queries.UpdateTaskDescription(ctx, ttssqlc.UpdateTaskDescriptionParams{
				Description: form.Description.Value,
				ID:          int64(params.ID),
			})

			if err != nil {
				return &api.Task{}, err
			}
		}

	}

	result := mapping.SqlTaskToApi(task)

	return &result, nil
}

func (s *TTApiService) GetTasksParentsTree(ctx context.Context, params api.GetTasksParentsTreeParams) (api.GetTasksParentsTreeRes, error) {
	session, _ := tsession.GetTSession(ctx, s.Queries)
	_, err := session.GetTask(int64(params.ID))
	if err == tsession.SessionForbiddenError {
		return ServiceErrors.Forbidden(), nil
	}
	if err == tsession.SessionUnauthorizedError {
		return ServiceErrors.Unauthorized(), nil
	}
	if err != nil {
		return &api.TaskRelationsList{}, err
	}

	rels, r_err := s.Queries.GetTaskParentsTree(ctx, int64(params.ID))

	if r_err != nil {
		return &api.TaskRelationsList{}, r_err
	}

	api_rels := []api.TaskRelation{}
	for _, dep := range rels {
		api_rels = append(api_rels, api.TaskRelation{
			TaskId:      int(dep.TaskID),
			DependsOnId: int(dep.DependsOnID),
		})
	}
	return &api.TaskRelationsList{Dependencies: api_rels}, nil
}

func (s *TTApiService) GetTasksParents(ctx context.Context, params api.GetTasksParentsParams) (api.GetTasksParentsRes, error) {
	session, _ := tsession.GetTSession(ctx, s.Queries)
	_, err := session.GetTask(int64(params.ID))
	if err == tsession.SessionForbiddenError {
		return ServiceErrors.Forbidden(), nil
	}
	if err == tsession.SessionUnauthorizedError {
		return ServiceErrors.Unauthorized(), nil
	}
	if err != nil {
		return &api.TaskDependencies{}, err
	}

	// FIXME
	rels, r_err := s.Queries.GetTaskParents(ctx, int64(params.ID))

	if r_err != nil {
		return &api.TaskDependencies{}, r_err
	}

	apiTasks := []api.Task{}

	for _, task := range rels {
		apiTasks = append(apiTasks, mapping.SqlTaskToApi(task))
	}

	return &api.TaskDependencies{
		Dependencies: apiTasks,
	}, nil
}

func (s *TTApiService) LinkTask(ctx context.Context, req api.OptDependencyForm, params api.LinkTaskParams) (api.LinkTaskRes, error) {
	// FIXME two users could link task at the same time
	// and cause circular dependency
	// - use mutex to lock simultaneos linking
	// - use after linking check if condition meets
	session, session_err := helpers.GetSession(ctx, *s.Queries)
	if session_err != nil {
		return ServiceErrors.Unauthorized(), nil
	}
	task, err := s.Queries.GetTask(ctx, int64(params.ID))
	if err != nil {
		return &api.TaskRelation{}, err
	}

	if task.CreatedBy != session.UserID {
		return ServiceErrors.Forbidden(), nil
	}

	parents, err_parents := s.Queries.GetTaskParentsTree(ctx, int64(params.ID))

	if err_parents != nil {
		return &api.TaskRelation{}, err_parents
	}

	d_task, d_err := s.Queries.GetTask(ctx, int64(req.Value.DependencyTaskId))

	if d_err != nil {
		return &api.TaskRelation{}, d_err
	}

	if d_task.CreatedBy != session.UserID {
		return ServiceErrors.Forbidden(), nil
	}

	for _, p := range parents {
		if p.TaskID == d_task.ID {
			return ServiceErrors.Circular(), nil
		}
	}

	rel, err_dep := s.Queries.CreateTaskDependency(ctx, ttssqlc.CreateTaskDependencyParams{
		TaskID:      task.ID,
		DependsOnID: d_task.ID,
		CreatedAt:   time.Now(),
	})

	if err_dep != nil {
		return &api.TaskRelation{}, err_dep
	}

	return &api.TaskRelation{
		TaskId:      int(rel.TaskID),
		DependsOnId: int(rel.DependsOnID),
	}, nil
}

func (s *TTApiService) SearchTask(ctx context.Context, params api.SearchTaskParams) (api.SearchTaskRes, error) {
	session, session_err := helpers.GetSession(ctx, *s.Queries)
	if session_err != nil {
		return ServiceErrors.Unauthorized(), nil
	}

	tasks, err := s.Queries.SearchTask(ctx, ttssqlc.SearchTaskParams{
		PlaintoTsquery: params.Title,
		CreatedBy:      session.UserID,
	})

	if err != nil {
		return &api.TasksList{}, err
	}

	r_tasks := []api.Task{}

	for _, t := range tasks {
		r_tasks = append(r_tasks, mapping.SqlTaskToApi(t))
	}

	return &api.TasksList{
		Tasks: r_tasks,
	}, nil
}

func (s *TTApiService) UnlinkTask(ctx context.Context, params api.UnlinkTaskParams) (api.UnlinkTaskRes, error) {
	session, _ := tsession.GetTSession(ctx, s.Queries)
	_, err := session.GetTask(int64(params.ID))
	if err == tsession.SessionForbiddenError {
		return ServiceErrors.Forbidden(), nil
	}
	if err == tsession.SessionUnauthorizedError {
		return ServiceErrors.Unauthorized(), nil
	}
	if err != nil {
		return &api.ResponseOk{}, err
	}

	unlink_err := s.Queries.UnlinkTask(ctx, ttssqlc.UnlinkTaskParams{
		TaskID:      int64(params.ID),
		DependsOnID: int64(params.DependsOnID),
	})

	// FIXME: recursively cleanup tasks

	return &api.ResponseOk{}, unlink_err
}

func (s *TTApiService) SetTaskDone(ctx context.Context, params api.SetTaskDoneParams) (api.SetTaskDoneRes, error) {
	session, _ := tsession.GetTSession(ctx, s.Queries)
	_, err := session.GetTask(int64(params.ID))
	if err == tsession.SessionForbiddenError {
		return ServiceErrors.Forbidden(), nil
	}
	if err == tsession.SessionUnauthorizedError {
		return ServiceErrors.Unauthorized(), nil
	}
	if err != nil {
		return &api.Task{}, err
	}

	updated_task, update_err := s.Queries.SetTaskDone(ctx, ttssqlc.SetTaskDoneParams{
		ID:     int64(params.ID),
		DoneAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	res_task := mapping.SqlTaskToApi(updated_task)

	return &res_task, update_err
}

func (s *TTApiService) SetTaskNotDone(ctx context.Context, params api.SetTaskNotDoneParams) (api.SetTaskNotDoneRes, error) {
	session, _ := tsession.GetTSession(ctx, s.Queries)
	_, err := session.GetTask(int64(params.ID))
	if err == tsession.SessionForbiddenError {
		return ServiceErrors.Forbidden(), nil
	}
	if err == tsession.SessionUnauthorizedError {
		return ServiceErrors.Unauthorized(), nil
	}
	if err != nil {
		return &api.Task{}, err
	}

	updated_task, update_err := s.Queries.SetTaskDone(ctx, ttssqlc.SetTaskDoneParams{
		ID:     int64(params.ID),
		DoneAt: sql.NullTime{Valid: false},
	})

	res_task := mapping.SqlTaskToApi(updated_task)

	return &res_task, update_err
}

// type TaskShortList = []api.TaskShort{}

func (s *TTApiService) GetTodo(ctx context.Context, params api.GetTodoParams) (api.GetTodoRes, error) {
	session, _ := tsession.GetTSession(ctx, s.Queries)
	task, err := session.GetTask(int64(params.ID))
	if err == tsession.SessionForbiddenError {
		return ServiceErrors.Forbidden(), nil
	}
	if err == tsession.SessionUnauthorizedError {
		return ServiceErrors.Unauthorized(), nil
	}
	if err != nil {
		return &api.TaskTodo{}, err
	}

	items, t_err := s.Queries.GetTaskDependenciesTree(ctx, int64(params.ID))

	if t_err != nil {
		return &api.TaskTodo{}, t_err
	}

	tasksMap := map[int64][]api.TaskShort{}

	for _, item := range items {
		id := item.ParentID
		taskShort := api.TaskShort{
			ID:     int(item.TaskID),
			Title:  item.Title,
			DoneAt: helpers.NilableTimeToApi(item.DoneAt),
		}
		tasksMap[id] = append(tasksMap[id], taskShort)
	}

	path := make([]api.TaskShort, 0)

	cursor := api.TaskShort{
		ID:     int(task.ID),
		Title:  task.Title,
		DoneAt: helpers.NilableTimeToApi(task.DoneAt),
	}
	path = append(path, cursor)
	cursor_ok := true

	for cursor_ok {
		deps, ok := tasksMap[int64(cursor.ID)]
		if ok {
			cursor = deps[0]
			path = append(path, cursor)
		} else {
			cursor_ok = false
		}
	}

	return &api.TaskTodo{
		Path: path,
	}, nil
}
