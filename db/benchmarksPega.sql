-- public.benchmarks definition

-- Drop table

-- DROP TABLE public.benchmarks;

CREATE TABLE public.benchmarks (
	id serial4 NOT NULL,
	bname text NULL,
	btype text NULL,
	kind text NULL,
	enabled bool DEFAULT false NOT NULL,
	branches _text NULL,
	priority int4 DEFAULT 4 NOT NULL,
	bcycle int4 DEFAULT 12 NOT NULL,
	boffset int4 DEFAULT 10 NOT NULL,
	burl text NULL,
	notification_email text NULL,
	issues text NULL,
	notes text NULL,
	"id""|""bname""|""btype""|""kind""|""enabled""|""branches""|""priority""|""bcy" varchar(256) NULL,
	CONSTRAINT benchmarks_pkey PRIMARY KEY (id)
);

truncate benchmarks;
select * from public."benchmarks" ba ;

