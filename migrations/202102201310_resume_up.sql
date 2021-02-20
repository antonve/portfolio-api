CREATE TABLE resume (
  slug varchar(10) NOT NULL,
  body jsonb NOT NULL,
  enabled boolean DEFAULT FALSE,
  PRIMARY KEY (slug),
  UNIQUE(slug)
);

CREATE SEQUENCE resume_logs_seq;

CREATE TABLE resume_logs (
  id bigint check (id > 0) NOT NULL DEFAULT NEXTVAL ('resume_logs_seq'),
  slug varchar(10) NOT NULL,
  created_at timestamp NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
  ip_address varchar(45) NOT NULL,
  user_agent text NOT NULL,
  PRIMARY KEY (id)
);

ALTER SEQUENCE resume_logs_seq RESTART WITH 1;
