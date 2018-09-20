-- Insert / Crud
INSERT INTO novel ("title", "author", "created") VALUES ('허균전', '홍길동');
INSERT INTO novel ("title", "author", "created") VALUES ('허생전', '박지원');
INSERT INTO novel ("title", "author", "created") VALUES ('구운몽', '김만중');
INSERT INTO novel ("title", "author", "created") VALUES ('제인 에어', '샬럿 브론테');
INSERT INTO novel ("title", "author", "created") VALUES ('폭풍의 언덕', '엘리스 벨(에밀리 브론테)');

-- Select / cRud
SELECT * FROM "novel" ORDER BY "_id" ASC LIMIT 100;
SELECT * FROM "novel" ORDER BY "_id" DESC LIMIT 100;

-- Update / crUd
UPDATE "novel" SET "title"='홍길동전', "author"='허균', WHERE "title"='허균전';

-- Delete / cruD
Delete FROM "novel" WHERE "_id"=1;
Delete FROM "novel" WHERE "author"='김만중';

-- Find
SELECT * FROM "novel" WHERE "title" LIKE '%전%' ORDER BY "_id" ASC LIMIT 100;
SELECT * FROM "novel" WHERE "author" LIKE '%브론테%' ORDER BY "_id" ASC LIMIT 100;
