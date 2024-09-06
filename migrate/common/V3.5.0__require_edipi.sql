DELETE FROM profile_project_roles WHERE profile_id IN (
  SELECT id FROM profile WHERE edipi IS NULL
);

DELETE FROM profile WHERE edipi IS NULL;

ALTER TABLE profile ALTER COLUMN edipi SET NOT NULL;
