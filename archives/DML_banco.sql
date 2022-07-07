CREATE TABLE public.books (
	id bigserial NOT NULL,
	name text NULL,
	description text NULL,
	medium_price numeric NULL,
	author text NULL,
	image_url text NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	CONSTRAINT books_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_books_deleted_at ON public.books USING btree (deleted_at);