-- https://astaxie.gitbooks.io/build-web-application-with-golang/en/05.4.html
-- Postgresql - Table creation
-- CREATE TABLE IF NOT EXISTS "books"
CREATE TABLE IF NOT EXISTS "novel"
(
	"_id" serial NOT NULL,
	"title" character varying(255) NOT NULL,
	"author" character varying(255) NOT NULL,
	-- "created" date,
	-- created_at timestamp with time zone DEFAULT current_timestamp,
	CONSTRAINT userinfo_pkey PRIMARY KEY ("_id")
)
-- ) WITH (OIDS=FALSE); // Not work with CockroachDB

-- Drop table
-- DROP TABLE [ IF EXISTS ] 이름 [, ...] [ CASCADE | RESTRICT ]
-- CASCADE 모든 의존 객체를 같이 삭제
-- RESTRICT 의존 객체가 있으면 작업 중지. 기본값.
DROP TABLE IF EXISTS "novel"
