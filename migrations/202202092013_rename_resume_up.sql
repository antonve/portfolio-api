ALTER TABLE resume
RENAME COLUMN enabled TO is_visible;

CREATE TABLE resume_visits (
  uuid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  slug varchar(10) NOT NULL,
  created_at timestamp NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
  ip_address varchar(45) NOT NULL,
  user_agent text NOT NULL
);