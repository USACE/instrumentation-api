\set ON_ERROR_STOP 1
begin;

\i '/sql/10-tables.sql'
\i '/sql/11-tables_aware.sql'
\i '/sql/15-views.sql'
\i '/sql/20-roles.sql'
\i '/sql/21-roles_aware.sql'
\i '/sql/30-data_hhd.sql'
\i '/sql/31-data_lrb.sql'
\i '/sql/32-data_lrc.sql'
\i '/sql/33-data_nae.sql'
\i '/sql/34-data_poa.sql'
\i '/sql/35-data_lrc_images.sql'

commit;