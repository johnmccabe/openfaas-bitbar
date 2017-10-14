package assets

// Logo base64 encoded OpenFaaS logo
const Logo string = "iVBORw0KGgoAAAANSUhEUgAAARcAAAA1CAYAAABrwK9dAAAABGdBTUEAALGPC/xhBQAAAAlwSFlzAAAN1wAADdcBQiibeAAABCNpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6dGlmZj0iaHR0cDovL25zLmFkb2JlLmNvbS90aWZmLzEuMC8iCiAgICAgICAgICAgIHhtbG5zOmV4aWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20vZXhpZi8xLjAvIgogICAgICAgICAgICB4bWxuczpkYz0iaHR0cDovL3B1cmwub3JnL2RjL2VsZW1lbnRzLzEuMS8iCiAgICAgICAgICAgIHhtbG5zOnhtcD0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLyI+CiAgICAgICAgIDx0aWZmOlJlc29sdXRpb25Vbml0PjI8L3RpZmY6UmVzb2x1dGlvblVuaXQ+CiAgICAgICAgIDx0aWZmOkNvbXByZXNzaW9uPjU8L3RpZmY6Q29tcHJlc3Npb24+CiAgICAgICAgIDx0aWZmOlhSZXNvbHV0aW9uPjkwPC90aWZmOlhSZXNvbHV0aW9uPgogICAgICAgICA8dGlmZjpPcmllbnRhdGlvbj4xPC90aWZmOk9yaWVudGF0aW9uPgogICAgICAgICA8dGlmZjpZUmVzb2x1dGlvbj45MDwvdGlmZjpZUmVzb2x1dGlvbj4KICAgICAgICAgPGV4aWY6UGl4ZWxYRGltZW5zaW9uPjI3OTwvZXhpZjpQaXhlbFhEaW1lbnNpb24+CiAgICAgICAgIDxleGlmOkNvbG9yU3BhY2U+MTwvZXhpZjpDb2xvclNwYWNlPgogICAgICAgICA8ZXhpZjpQaXhlbFlEaW1lbnNpb24+NTM8L2V4aWY6UGl4ZWxZRGltZW5zaW9uPgogICAg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"

// MonoLogo is a base64 representation of a monochrome Logo
const MonoLogo string = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAQAAADZc7J/AAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAAmJLR0QA/4ePzL8AAAAJcEhZcwAAFiUAABYlAUlSJPAAAAAHdElNRQfhCgoALzQHaVaaAAAAXnRFWHRSYXcgcHJvZmlsZSB0eXBlIGlwdGMACmlwdGMKICAgICAgMjgKMzg0MjQ5NGQwNDA0MDAwMDAwMDAwMDBmMWMwMjZlMDAwMzUyNDY0NzFjMDIwMDAwMDIwMDA0MDAKYJo9ngAAArRJREFUSMeN1U1oXFUUB/Dfm4S8SDWtKdSSYKGi1VBcWIMBdwFxUYVu4sKNCC6k3YhriTupG1eC4MdWXaRdlIqoCEZEsDWlizLGSKmpxY8gtZlMp8nMezPXxcy8vJm8tPMOD+6599z/+bj/c68gEJu3qCIMKBWL5sVBEAXRlDOu+tCSDYN9Y6a97lFzYZlY2alBXefFKWVxZN50ONHvItoTavcOJDpniUXHe3Cf8rYl1+/ieSQbHbdIxYGOOu0DNzK7p3s2veecH5T9o+HrbPaAyrCxrHTHnEQidR8uRMtWOlL2oiMgYDzLYcMYwWgHb0Irc9nqifqkizntahbBqFDaLkn4y6+Z0sCGmlQDK9ZztRvPF7KUV3yfjWLBmD2GjWDFTdQ1tLTsjUq7AXyM1Hfe9KNGNrsZ/rSK2IiS4Kb785uC0cB+Q0Fgzr4gcCmX8+XAE5437bC9oiDwULcGXYBvXXao5+D2mfGKd5x1uo8Hkbc0PdkGGO7EsWpWOfrMl647ZNyDqj4PFzqMmzFr3X/W1DznVUdEbvSm8IaaoKqqbl1V3S23vRB4wJJUS0sikUoEQbU/hWc0Cli75YRLfZxoyzd9AIHlArNm5++X1OxOgJdtDdzJy9tMzFP508I0dsqmw8UAo84PAFHzbL4XcgCdW2ZNbZetLakvPJyz3gkQGDLnK39r2swVM/Gbd032UaoIIFs86kqnrFteMlxoczeAwJxb3V7YxSJH5TZlhzzuEdeshCYeE4Op9hXbt9rbjYEJC1KpmlTiEwf9ktHmNRMWJBL/2lT3kYPbKVRMOup9SQ9l79jK6WudDuhKQ2LBjEkVrmiqF/I97aF00bG2lEvOS4yICt6NoV3uraxkWn4mdk194C7IS93vYoEpfxSmcC9Kr5oK2ucZO+2i2wNvvuOn7vP+P4/dyPe/cPxbAAAAJXRFWHRkYXRlOmNyZWF0ZQAyMDE3LTEwLTE0VDEwOjIzOjA1KzAxOjAw7YB2ewAAACV0RVh0ZGF0ZTptb2RpZnkAMjAxNy0xMC0wOVQyMzo0Nzo1MiswMTowML+DVBsAAAAZdEVYdFNvZnR3YXJlAHd3dy5pbmtzY2FwZS5vcmeb7jwaAAAAG3RFWHRUaXRsZQBPcGVuRmFhU19sb2dvX3N0YWNrZWRx2Jn2AAAAAElFTkSuQmCC"

// BlueLogo is a base64 representation of a monochrome blue log - #067FD1
const BlueLogo string = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAABGdBTUEAALGPC/xhBQAAAAlwSFlzAAAWJQAAFiUBSVIk8AAABCRpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6dGlmZj0iaHR0cDovL25zLmFkb2JlLmNvbS90aWZmLzEuMC8iCiAgICAgICAgICAgIHhtbG5zOmV4aWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20vZXhpZi8xLjAvIgogICAgICAgICAgICB4bWxuczpkYz0iaHR0cDovL3B1cmwub3JnL2RjL2VsZW1lbnRzLzEuMS8iCiAgICAgICAgICAgIHhtbG5zOnhtcD0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLyI+CiAgICAgICAgIDx0aWZmOlJlc29sdXRpb25Vbml0PjI8L3RpZmY6UmVzb2x1dGlvblVuaXQ+CiAgICAgICAgIDx0aWZmOkNvbXByZXNzaW9uPjU8L3RpZmY6Q29tcHJlc3Npb24+CiAgICAgICAgIDx0aWZmOlhSZXNvbHV0aW9uPjE0NDwvdGlmZjpYUmVzb2x1dGlvbj4KICAgICAgICAgPHRpZmY6T3JpZW50YXRpb24+MTwvdGlmZjpPcmllbnRhdGlvbj4KICAgICAgICAgPHRpZmY6WVJlc29sdXRpb24+MTQ0PC90aWZmOllSZXNvbHV0aW9uPgogICAgICAgICA8ZXhpZjpQaXhlbFhEaW1lbnNpb24+MzI8L2V4aWY6UGl4ZWxYRGltZW5zaW9uPgogICAgICAgICA8ZXhpZjpDb2xvclNwYWNlPjE8L2V4aWY6Q29sb3JTcGFjZT4KICAgICAgICAgPGV4aWY6UGl4ZWxZRGltZW5zaW9uPjMyPC9leGlmOlBpeGVsWURpbWVuc2lvbj4KICAg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"

// RedLogo is a base64 representation of a monochrome red logo - #DA0002
const RedLogo string = "iVBORw0KGgoAAAANSUhEUgAAACAAAAAgCAYAAABzenr0AAAABGdBTUEAALGPC/xhBQAAAAlwSFlzAAAWJQAAFiUBSVIk8AAABCRpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IlhNUCBDb3JlIDUuNC4wIj4KICAgPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICAgICAgPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIKICAgICAgICAgICAgeG1sbnM6dGlmZj0iaHR0cDovL25zLmFkb2JlLmNvbS90aWZmLzEuMC8iCiAgICAgICAgICAgIHhtbG5zOmV4aWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20vZXhpZi8xLjAvIgogICAgICAgICAgICB4bWxuczpkYz0iaHR0cDovL3B1cmwub3JnL2RjL2VsZW1lbnRzLzEuMS8iCiAgICAgICAgICAgIHhtbG5zOnhtcD0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLyI+CiAgICAgICAgIDx0aWZmOlJlc29sdXRpb25Vbml0PjI8L3RpZmY6UmVzb2x1dGlvblVuaXQ+CiAgICAgICAgIDx0aWZmOkNvbXByZXNzaW9uPjU8L3RpZmY6Q29tcHJlc3Npb24+CiAgICAgICAgIDx0aWZmOlhSZXNvbHV0aW9uPjE0NDwvdGlmZjpYUmVzb2x1dGlvbj4KICAgICAgICAgPHRpZmY6T3JpZW50YXRpb24+MTwvdGlmZjpPcmllbnRhdGlvbj4KICAgICAgICAgPHRpZmY6WVJlc29sdXRpb24+MTQ0PC90aWZmOllSZXNvbHV0aW9uPgogICAgICAgICA8ZXhpZjpQaXhlbFhEaW1lbnNpb24+MzI8L2V4aWY6UGl4ZWxYRGltZW5zaW9uPgogICAgICAgICA8ZXhpZjpDb2xvclNwYWNlPjE8L2V4aWY6Q29sb3JTcGFjZT4KICAgICAgICAgPGV4aWY6UGl4ZWxZRGltZW5zaW9uPjMyPC9leGlmOlBpeGVsWURpbWVuc2lvbj4KICAg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"
