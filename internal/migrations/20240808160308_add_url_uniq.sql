-- +goose Up
-- +goose StatementBegin
CREATE UNIQUE INDEX url_list_url_idx ON public.url_list USING btree (url);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS url_list_url_idx;
-- +goose StatementEnd
