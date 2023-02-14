-- Offset descriptor
CREATE TABLE IF NOT EXISTS offset_descriptor (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT
);

INSERT INTO offset_descriptor (id, name) VALUES
    ('a0834daf-7c94-4ed9-acb9-5e6ede35e9c1', '- / +'),
    ('9809dcbe-45a5-4914-84ca-9d326c057d34', 'DS / US'),
    ('83f99462-d774-4867-8777-6ca55e0c8732', 'L / R');

ALTER TABLE instrument ALTER COLUMN station_offset TYPE DOUBLE PRECISION;
ALTER TABLE instrument ADD COLUMN offset_descriptor_id UUID REFERENCES offset_descriptor (id);
ALTER TABLE instrument SET offset_descriptor_id = 'a0834daf-7c94-4ed9-acb9-5e6ede35e9c1'::UUID;

ALTER TABLE instrument
ALTER COLUMN offset_descriptor_id SET NOT NULL,
ADD CONSTRAINT offset_descriptor_id_fkey
FOREIGN KEY (offset_descriptor_id) REFERENCES offset_descriptor (id);
