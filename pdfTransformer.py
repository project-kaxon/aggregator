import json
import sys
from logging import error
from pathlib import Path
from pprint import pprint

import fitz

if len(sys.argv) < 2:
	print('Error: Must supply command arguments', file=sys.stderr)
	exit(1);

arr = []
for pdf_file in sys.argv[1:]:
	pdf_file = Path(pdf_file)
	doc = fitz.open(pdf_file.resolve())
	toc = doc.get_toc(simple=False)

	for level, title, hierarchy, dest in toc:
		title = title.strip()

		offset = 0
		if dest['kind'] == 1:
			offset = dest["to"].y

		hierarchy = title[:title.find(' ')]
		if '.' in hierarchy:
			arr.append({
				'indexes': {
					'hierarchy': hierarchy,
					'title': title[title.find(' ')+1:],
				},
				'title': title,
				'page': dest['page'],
				'offset': offset
			})

	json_text = json.dumps(arr)
	json_file = pdf_file.parent / f'{pdf_file.stem}.json'
	Path(json_file).write_text(json_text);
