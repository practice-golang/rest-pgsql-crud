-- Source: http://freeprog.tistory.com/248
-- 역슬래쉬 명령은 여기서 안 된다.

-- tablespace 만들기
CREATE TABLESPACE testdbtablespace LOCATION 'c:\dbdata\testdb_ts';

--  tablespace 조회 명령어 
\db;

-- db 생성1 / tablespace 지정 안하는 경우
CREATE DATABASE testdb;

-- db 생성2 / tablespace 지정
CREATE DATABASE testdb TABLESPACE testdbtablespace;

-- db 조회
\l

-- db 연결
\c dbname

-- schema 목록
\dn

-- table 생성
CREATE TABLE teacher (
    id          serial PRIMARY KEY,
    name        varchar(40),
    jumin       char(13),
    adm         date,
    retire      timestamptz
);

-- table 목록 출력
\dt

-- table 구조 출력
\d tablename

-- Insert / Crud
INSERT INTO teacher (name, jumin, special, adm, retire)
VALUES ('홍길동', '5501011234567', false, '2012-01-10', NULL);
-- 여러 줄 한 번에 입력
INSERT INTO teacher (name, jumin, special, adm, retire) VALUES
('이태백', '4501111234577', true, '1998-11-10', '2012-01-30 18:30:00'),
('황진희', '5601222234577', false, '1997-05-10', NULL),
('김삿갓', '3601021235577', false, '1999-04-11', NULL);

-- Select / crUd
SELECT * from teacher;

-- Update / crUd
UPDATE teacher SET retire='2015-12-31 21:00:30'
WHERE name='김삿갓';

-- Delete / crUd
DELETE FROM teacher WHERE id=3;

