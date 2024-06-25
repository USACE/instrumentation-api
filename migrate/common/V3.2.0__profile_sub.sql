-- As CAC-Only users log in, their "sub" will be added to profile table
-- When non-CAC user access is added, the sub will be required to lookup profile
ALTER TABLE profile ADD COLUMN sub text;
