DROP TABLE resume_visits CASCADE;

ALTER TABLE resume
RENAME COLUMN is_visible TO enabled;