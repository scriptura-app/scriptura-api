SELECT verse.verse_num, bible_es_rvr1960.text
FROM bible_es_rvr1960
JOIN verse
ON bible_es_rvr1960.verse_id = verse.id
WHERE verse.book_id = 4 AND
verse.chapter_num = 2 AND 
verse.verse_num >= 3 AND
verse.verse_num <= 20;