SELECT
	v.id,
	ber."text",
	coalesce(
		(
			SELECT
				json_agg(to_json(p))
			FROM
				people p
				JOIN people_verses pv ON p.id = pv.person_id
			WHERE
				v.id = pv.verse_id
		),
		'[]' :: json
	) people,
	coalesce(
		(
			SELECT
				json_agg(to_json(p2))
			FROM
				places p2
				JOIN places_verses pv2 ON p2.id = pv2.place_id
			WHERE
				v.id = pv2.verse_id
		),
		'[]' :: json
	) places,
	coalesce(
		(
			SELECT
				json_agg(to_json(e))
			FROM
				EVENTS e
				JOIN events_verses ev ON e.id = ev.event_id
			WHERE
				v.id = ev.verse_id
		),
		'[]' :: json
	) EVENTS
FROM
	verses v
	JOIN bible_es_rv1960 ber ON v.id = ber.verse_id
WHERE
	v.book_id = 1
	AND v.chapter_num = 4
	AND v.verse_num BETWEEN 1
	AND 20;