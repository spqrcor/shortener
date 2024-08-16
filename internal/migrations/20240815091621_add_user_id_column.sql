-- +goose Up
-- +goose StatementBegin
ALTER TABLE public.url_list ADD COLUMN IF NOT EXISTS user_id uuid;
CREATE INDEX url_list_user_id_idx ON public.url_list USING btree (user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS url_list_user_id_idx;
ALTER TABLE public.url_list DROP COLUMN IF EXISTS user_id;
-- +goose StatementEnd