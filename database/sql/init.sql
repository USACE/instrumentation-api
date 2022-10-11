begin;

\i '/sql/10-tables.sql'
\i '/sql/11-tables_aware.sql'
\i '/sql/15-views.sql'
\i '/sql/16-functions-triggers.sql'
\i '/sql/20-roles.sql'
\i '/sql/21-roles_aware.sql'
\i '/sql/29-data_usace_projects.sql'
\i '/sql/30-data_hhd.sql'
\i '/sql/32-data_lrc.sql'
\i '/sql/33-data_nae.sql'
\i '/sql/34-data_poa.sql'
\i '/sql/35-data_lrc_images.sql'
\i '/sql/36-data_c44.sql'
\i '/sql/37-data_projects_feb2022.sql'
-- New district streamgages should be appended below and not in 
-- alphabetical order due to script that ensures slugs are unique.
\i '/sql/district_streamgages/lrh.sql'
\i '/sql/district_streamgages/lrn.sql'
\i '/sql/district_streamgages/mvk.sql'
\i '/sql/district_streamgages/mvn.sql'
\i '/sql/district_streamgages/nwdm.sql'
\i '/sql/district_streamgages/sas.sql'
\i '/sql/district_streamgages/nae.sql'
\i '/sql/district_streamgages/spa.sql'
\i '/sql/district_streamgages/lrb.sql'
\i '/sql/district_streamgages/lrc.sql'
commit;