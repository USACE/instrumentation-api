CREATE TABLE division (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    name TEXT,
    initials VARCHAR(3)
);

CREATE TABLE office (
    id UUID PRIMARY KEY NOT NULL
);

CREATE TABLE district (
    id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    division_id UUID NOT NULL REFERENCES division (id),
    name TEXT,
    initials VARCHAR(3),
    office_id UUID REFERENCES office (id)
);

CREATE UNIQUE INDEX unique_district_office_id ON district (office_id)
WHERE office_id IS NOT NULL;

INSERT INTO office (id) VALUES
    ('00000000-0000-0000-0000-000000000000'),
    ('17fa25b8-44a0-4e6d-9679-bdf6b0ee6b1a'),
    ('d8f8934d-e414-499d-bd51-bc93bbde6345'),
    ('a8192ad1-206c-4da6-b19e-b7ba7a67aa1f'),
    ('2f160ba7-fd5f-4716-8ced-4a29f75065a6'),
    ('433a554d-7b27-4046-89eb-906788eb4046'),
    ('552e59f7-c0cc-4689-8a4d-e791c028430a'),
    ('61291eaf-d62f-4846-ad95-87cc86b56851'),
    ('1245e3c0-fc72-4621-86b2-24ff7de21f88'),
    ('f81f5659-ce57-4c87-9c7a-0d685a983bfd'),
    ('81557734-7046-4c55-90ac-066dd882166a'),
    ('565be474-0c68-44a6-8e66-b833a39685bd'),
    ('2cf60156-f22a-418a-bc9f-a28960ad0321'),
    ('b9cca282-eb91-4ea1-b075-d067b4420184'),
    ('7ed4821f-9e37-4c56-8baf-05c1b5bcc84c'),
    ('30cb06ee-bd94-4c49-a945-e92735e7bdc1'),
    ('a47f1ef4-1017-43c1-bf36-67f57376d163'),
    ('1989e3fc-f12a-42da-a263-c3ae978e2c09'),
    ('5b35ea7c-8d1b-481a-956b-b32939093db4'),
    ('665ffb00-ccba-488c-83c5-2083543cacd7'),
    ('8b0a732d-d511-4332-b2e7-dd6943a597e9'),
    ('007cbff5-6946-4b9b-a3f7-0bef4406f122'),
    ('30266178-d32a-4e07-aea1-1f7182ed245e'),
    ('d4501358-1c48-45cb-86f3-f1a31e9bd93f'),
    ('b4f45596-70e5-4a12-a894-a64300648244'),
    ('9ffc189c-ad40-4fbf-bc06-2098c6cb920e'),
    ('0154184e-2509-4485-b449-8eff4ab52eef'),
    ('ba1f7846-43d9-4a21-9876-27c59510d9c0'),
    ('11b5fe49-fe36-4a06-a0da-d55b1b62b1fb'),
    ('b8cec5bc-f975-4bed-993d-8f913ca51673'),
    ('ff52a84b-356a-4173-a8df-89a1b408d354'),
    ('cf9b1f4d-1cd3-4a00-b73d-b6f8fe75915e'),
    ('f3f0d7ff-19b6-4167-a3f1-5c04f5a0fe4d'),
    ('72ee5695-cdaa-4182-b0c1-4d27f1a3f570'),
    ('131daa5c-49c2-4488-be6b-bd638a83a03f'),
    ('fe29f6e2-e200-44a4-9545-a4680ab9366e');

INSERT INTO division (id, name, initials) VALUES
    ('3ab59912-d9b9-476b-be27-60ef1638f041', 'Great Lakes and Ohio River Division', 'LRD'),
    ('efdc137c-3077-40aa-81ab-172aea45f6d6', 'Mississippi Valley Division', 'MVD'),
    ('fcf3c8bc-3e96-4e27-acbd-9956da2e6616', 'North Atlantic Division', 'NAD'),
    ('21ea5cf4-9734-41dc-ab40-f5472bd5bbfd', 'Northwestern Division', 'NWD'),
    ('ef2760a7-1d6f-4c12-b1ac-0f0b1cad2b7d', 'Pacific Ocean Division', 'POD'),
    ('b23031fc-60f6-4ae4-a82c-f7a6210f8614', 'South Atlantic Division', 'SAD'),
    ('f3addd12-1f7d-43a0-b692-c42be934ead9', 'South Pacific Division', 'SPD'),
    ('b312c218-5c69-4a17-a7fb-f61f2b86b566', 'Southwestern Division', 'SWD'),
    ('0634eca3-c13e-41de-b012-8c456d6422d7', 'Transatlantic Division', 'TAD');

INSERT INTO district (id, division_id, name, initials, office_id) VALUES
    ('321cc3d1-99e9-4b14-850e-7564afd47a79', '3ab59912-d9b9-476b-be27-60ef1638f041', 'Buffalo District', 'LRB', '17fa25b8-44a0-4e6d-9679-bdf6b0ee6b1a'),
    ('7d7636a0-9607-46d0-b3d1-89b19e2efded', '3ab59912-d9b9-476b-be27-60ef1638f041', 'Chicago District', 'LRC', 'd8f8934d-e414-499d-bd51-bc93bbde6345'),
    ('e12d168d-374a-49de-8940-a65ef93cb8d2', '3ab59912-d9b9-476b-be27-60ef1638f041', 'Detroit District', 'LRE', 'a8192ad1-206c-4da6-b19e-b7ba7a67aa1f'),
    ('2fefd988-9d42-4289-b21a-753756ee8748', '3ab59912-d9b9-476b-be27-60ef1638f041', 'Huntington District','LRH', '2f160ba7-fd5f-4716-8ced-4a29f75065a6'),
    ('cb26f34a-77c7-442c-8e46-4a9b4864266e', '3ab59912-d9b9-476b-be27-60ef1638f041', 'Louisville District','LRL', '433a554d-7b27-4046-89eb-906788eb4046'),
    ('be865d1c-8348-44b1-a2b2-accf959ed41d', '3ab59912-d9b9-476b-be27-60ef1638f041', 'Nashville District', 'LRN', '552e59f7-c0cc-4689-8a4d-e791c028430a'),
    ('02c990a5-b4f8-43ac-84d9-2f6ac753813e', '3ab59912-d9b9-476b-be27-60ef1638f041', 'Pittsburgh District','LRP', '61291eaf-d62f-4846-ad95-87cc86b56851'),
    ('9f6b422a-39cf-4d75-b23a-39b2f450c354', 'efdc137c-3077-40aa-81ab-172aea45f6d6', 'Memphis District', 'MVM', '1245e3c0-fc72-4621-86b2-24ff7de21f88'),
    ('8450b1c6-e71e-4e7c-89a9-4b5b198d6595', 'efdc137c-3077-40aa-81ab-172aea45f6d6', 'New Orleans District','MVN', 'f81f5659-ce57-4c87-9c7a-0d685a983bfd'),
    ('e716e041-fa2b-4dc9-bd3c-cbde5912af3d', 'efdc137c-3077-40aa-81ab-172aea45f6d6', 'Rock Island District','MVR', '81557734-7046-4c55-90ac-066dd882166a'),
    ('43b884d2-e91e-4f2d-b01b-4e1ba109ca2d', 'efdc137c-3077-40aa-81ab-172aea45f6d6', 'St. Louis District', 'MVS', '565be474-0c68-44a6-8e66-b833a39685bd'),
    ('487f09a1-a68b-46cd-8310-57ae5cecdf8d', 'efdc137c-3077-40aa-81ab-172aea45f6d6', 'St. Paul District', 'MVP', '2cf60156-f22a-418a-bc9f-a28960ad0321'),
    ('e8e87a67-9c8f-4a6b-8345-866b5845b910', 'efdc137c-3077-40aa-81ab-172aea45f6d6', 'Vicksburg District', 'MVK', 'b9cca282-eb91-4ea1-b075-d067b4420184'),
    ('7affcf88-6ee6-4fc5-aa09-32594b334dbc', 'fcf3c8bc-3e96-4e27-acbd-9956da2e6616', 'Baltimore District', 'NAB', '7ed4821f-9e37-4c56-8baf-05c1b5bcc84c'),
    ('e708078e-1c9c-408a-bd13-82bbf5fcb3ac', 'fcf3c8bc-3e96-4e27-acbd-9956da2e6616', 'Europe District', 'NAU', NULL),
    ('41219a9a-d689-4bef-9aef-bea23b17ea8d', 'fcf3c8bc-3e96-4e27-acbd-9956da2e6616', 'New England District','NAE', '30cb06ee-bd94-4c49-a945-e92735e7bdc1'),
    ('3315572f-b9b1-4e04-8594-cb02bc8e445e', 'fcf3c8bc-3e96-4e27-acbd-9956da2e6616', 'New York District', 'NAN', NULL),
    ('630eb743-876a-4286-8112-f0702de98bb2', 'fcf3c8bc-3e96-4e27-acbd-9956da2e6616', 'Norfolk District', 'NAO', 'a47f1ef4-1017-43c1-bf36-67f57376d163'),
    ('2f530073-b1b2-420b-8634-1680b4ab3c05', 'fcf3c8bc-3e96-4e27-acbd-9956da2e6616', 'Philadelphia District','NAP', '1989e3fc-f12a-42da-a263-c3ae978e2c09'),
    ('90c3e812-bdde-45ff-ae4f-f3392991d44e', '21ea5cf4-9734-41dc-ab40-f5472bd5bbfd', 'Kansas City District','NWK', '5b35ea7c-8d1b-481a-956b-b32939093db4'),
    ('218fc12f-d257-4580-b74f-5c70c1f04cce', '21ea5cf4-9734-41dc-ab40-f5472bd5bbfd', 'Omaha District',  'NOW', '665ffb00-ccba-488c-83c5-2083543cacd7'),
    ('5cf0df51-21d8-473e-b46c-edfa58a0a109', '21ea5cf4-9734-41dc-ab40-f5472bd5bbfd', 'Portland District', 'NWP', '8b0a732d-d511-4332-b2e7-dd6943a597e9'),
    ('db615b95-5a88-40f1-8cfe-d1187a0a66e7', '21ea5cf4-9734-41dc-ab40-f5472bd5bbfd', 'Seattle District', 'NWS', '007cbff5-6946-4b9b-a3f7-0bef4406f122'),
    ('0cb58fa0-a002-433c-897e-fa01d3edb399', '21ea5cf4-9734-41dc-ab40-f5472bd5bbfd', 'Walla Walla District','NWW', '30266178-d32a-4e07-aea1-1f7182ed245e'),
    ('a26d3e91-119c-4cc1-a5ce-9ac4283183bb', 'ef2760a7-1d6f-4c12-b1ac-0f0b1cad2b7d', 'Alaska District', 'POA', NULL),
    ('6a4eba67-2182-42e5-891d-179aac028829', 'ef2760a7-1d6f-4c12-b1ac-0f0b1cad2b7d', 'Far East District', 'POF', NULL),
    ('e37254c5-7c9c-4e72-ab4e-7f7d2558d656', 'ef2760a7-1d6f-4c12-b1ac-0f0b1cad2b7d', 'Honolulu District', 'POH', NULL),
    ('ca516ec6-4102-4d7b-8212-93e6a21868ad', 'ef2760a7-1d6f-4c12-b1ac-0f0b1cad2b7d', 'Japan Engineer District','POJ', NULL),
    ('d0dc5985-aa39-426c-9f95-fc70716ed96a', 'b23031fc-60f6-4ae4-a82c-f7a6210f8614', 'Charleston District', 'SAC', 'd4501358-1c48-45cb-86f3-f1a31e9bd93f'),
    ('c1bcd699-17cc-422e-9e53-d82f41aeaeb8', 'b23031fc-60f6-4ae4-a82c-f7a6210f8614', 'Jacksonville District','SAJ', 'b4f45596-70e5-4a12-a894-a64300648244'),
    ('be54ad0b-4c00-4983-b2fc-76bc341b0432', 'b23031fc-60f6-4ae4-a82c-f7a6210f8614', 'Mobile District', 'SAM', '9ffc189c-ad40-4fbf-bc06-2098c6cb920e'),
    ('c3cb72c5-a082-430a-b682-6e6765fea619', 'b23031fc-60f6-4ae4-a82c-f7a6210f8614', 'Savannah District', 'SAS', '0154184e-2509-4485-b449-8eff4ab52eef'),
    ('4d6b988b-5b70-41c6-bbb3-e3e580c17b6b', 'b23031fc-60f6-4ae4-a82c-f7a6210f8614', 'Wilmington District', 'SAW', 'ba1f7846-43d9-4a21-9876-27c59510d9c0'),
    ('58cc7a26-21dd-400b-b5ec-e84b8ae9f2eb', 'f3addd12-1f7d-43a0-b692-c42be934ead9', 'Albuquerque District', 'SPA', '11b5fe49-fe36-4a06-a0da-d55b1b62b1fb'),
    ('0386f35d-b2b7-446e-ab31-524034c3512e', 'f3addd12-1f7d-43a0-b692-c42be934ead9', 'Los Angeles District', 'SPL', 'b8cec5bc-f975-4bed-993d-8f913ca51673'),
    ('f3255f75-8ad9-414f-9e78-a6f054352b86', 'f3addd12-1f7d-43a0-b692-c42be934ead9', 'Sacramento District ', 'SPK', 'ff52a84b-356a-4173-a8df-89a1b408d354'),
    ('1970c46a-7b53-48ad-8749-4575b1df3ad1', 'f3addd12-1f7d-43a0-b692-c42be934ead9', 'San Francisco District', 'SPN', 'cf9b1f4d-1cd3-4a00-b73d-b6f8fe75915e'),
    ('09d7c0f5-a8ae-4869-b9f4-4e8dcfb56b5d', 'b312c218-5c69-4a17-a7fb-f61f2b86b566', 'Fort Worth District','SWF', 'f3f0d7ff-19b6-4167-a3f1-5c04f5a0fe4d'),
    ('a80c242e-07db-4fb0-a22d-342c2dd04d12', 'b312c218-5c69-4a17-a7fb-f61f2b86b566', 'Galveston District', 'SWG', '72ee5695-cdaa-4182-b0c1-4d27f1a3f570'),
    ('d30b429b-8f33-47b2-a276-bf86f8f50346', 'b312c218-5c69-4a17-a7fb-f61f2b86b566', 'Little Rock District','SWL', '131daa5c-49c2-4488-be6b-bd638a83a03f'),
    ('ddeec2a1-b416-4610-b24d-bf8065b07b79', 'b312c218-5c69-4a17-a7fb-f61f2b86b566', 'Tulsa District',  'SWT', 'fe29f6e2-e200-44a4-9545-a4680ab9366e'),
    ('c4806796-c244-4459-a49c-123425fad897', '0634eca3-c13e-41de-b012-8c456d6422d7', 'Middle East District', 'TAM', NULL),
    ('6e53b51b-c865-40af-af4a-63913ad2bb2d', '0634eca3-c13e-41de-b012-8c456d6422d7', 'Transatlantic Afghanistan District', 'TAD', NULL);

ALTER TABLE project ADD CONSTRAINT project_office_id_fkey
FOREIGN KEY (office_id) REFERENCES office (id);
