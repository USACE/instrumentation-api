CREATE OR REPLACE VIEW v_district_rollup AS (
    SELECT
        ac.alert_type_id                    AS alert_type_id,
        dt.office_id                        AS office_id,
        dt.initials                         AS district_initials,
        prj.name                            AS project_name,
        prj.id                              AS project_id,
        DATE_TRUNC('month', sub.due_date)   AS the_month,
        COUNT(sub.*)                        AS expected_total_submittals,
        COUNT(sub.completion_date) FILTER (
            WHERE sub.completion_date IS NOT NULL
        )                                   AS actual_total_submittals,
        COUNT(sub.*) FILTER (
            WHERE sub.submittal_status_id = '84a0f437-a20a-4ac2-8a5b-f8dc35e8489b'
        )                                   AS red_submittals,
        COUNT(sub.*) FILTER (
            WHERE sub.submittal_status_id = 'ef9a3235-f6e2-4e6c-92f6-760684308f7f'
        )                                   AS yellow_submittals,
        COUNT(sub.*) FILTER (
            WHERE sub.submittal_status_id = '0c0d6487-3f71-4121-8575-19514c7b9f03'
        )                                   AS green_submittals
    FROM alert_config ac
    INNER JOIN project prj ON ac.project_id = prj.id
    LEFT JOIN district dt ON dt.office_id = prj.office_id
    LEFT JOIN submittal sub ON sub.alert_config_id = ac.id
    WHERE DATE_TRUNC('month', sub.due_date) >= NOW() - INTERVAL '1 year' AND sub.due_date <= NOW()
    GROUP BY ac.alert_type_id, dt.office_id, dt.initials, prj.id, prj.name, DATE_TRUNC('month', sub.due_date)
	ORDER BY DATE_TRUNC('month', sub.due_date), ac.alert_type_id
);

GRANT SELECT ON
    v_district_rollup
TO instrumentation_reader;
