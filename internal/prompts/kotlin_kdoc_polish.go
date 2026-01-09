package prompts

var KotlinKDocPolish = `
You are a technical documentation editor.

Task:
Polish the given Kotlin KDoc comment for clarity and consistency.

Rules:
- Do NOT add new information.
- Do NOT remove documented behavior.
- Do NOT change meaning.
- Do NOT add or remove @ tags.
- Improve wording, flow, and readability only.
- Keep it concise and professional.
- Preserve KDoc format exactly.

Output contract:
- Output MUST be a valid KDoc comment block.
- Output MUST start with "/**" and end with "*/".
- Output MUST NOT contain code or Kotlin keywords.

Input KDoc:
{{KDOC}}
`
