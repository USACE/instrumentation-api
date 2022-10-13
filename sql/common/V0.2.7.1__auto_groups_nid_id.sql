-- INSERT INTO instrument_group (slug, name, description, project_id)
	(select 
	distinct lower(i.nid_id), 
	i.nid_id, 
	CONCAT(nid_id, ' automated group'),
	i.project_id
	from instrument i
	where i.nid_id is not null
	and i.project_id = 'b09015da-eae4-4f96-ad67-745ac0e3ea5b' 
	order by i.nid_id
	) on conflict do nothing;
	
-- INSERT INTO instrument_group_instruments(instrument_id, instrument_group_id)
   (select i.id, ig.id 
	from instrument i
   left join instrument_group ig on lower(ig.slug) = lower(i.nid_id)
	and ig.project_id = i.project_id
   where i.project_id = 'b09015da-eae4-4f96-ad67-745ac0e3ea5b'
   and i.id is not null 
   and ig.id is not null
   ) on conflict do nothing;