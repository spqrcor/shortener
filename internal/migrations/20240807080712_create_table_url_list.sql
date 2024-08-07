-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.url_list (
  short_url varchar(30) NOT NULL,
  url varchar NOT NULL,
  created_at timestamptz DEFAULT now() NOT NULL,
  id uuid DEFAULT gen_random_uuid() NOT NULL,
  CONSTRAINT url_list_pkey PRIMARY KEY (id)
);
CREATE INDEX url_list_short_url_idx ON public.url_list USING btree (short_url);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.url_list;
-- +goose StatementEnd
