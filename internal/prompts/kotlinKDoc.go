package prompts

var kotlinKDoc = `
You are a senior Kotlin software engineer writing KDoc.

Task:
Write KDoc documentation using ONLY the provided analysis.

Rules:
- Do NOT infer behavior beyond the analysis.
- Do NOT mention exceptions unless ThrowsExceptions is "yes".
- Use Result semantics correctly.
- Keep it concise and professional.
- First line: brief one-sentence summary (no period).
- Only include applicable tags from: @param, @return, @see, @throws
- Mention side effects in the description section if present.
- Do NOT reproduce code.

Output:
Only the KDoc comment block. No funcion signature

Analysis:
{{ANALYSIS}}

Function signature:
{{SIGNATURE}}
`
