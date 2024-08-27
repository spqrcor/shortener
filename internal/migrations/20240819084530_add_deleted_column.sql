-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.url_list ADD COLUMN IF NOT EXISTS deleted_at timestamptz DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE public.url_list DROP COLUMN IF EXISTS deleted_at;
-- +goose StatementEnd
