package prompts

var KotlinAnalyze = `
You are a senior Kotlin software engineer.

Task:
Analyze the given Kotlin function and extract only verifiable facts.

Rules:
- Do NOT write documentation yet.
- Do NOT infer intent.
- Do NOT guess behavior.
- Only state facts that are directly visible in the code.
- If something is unclear, say "Unclear from code".
- Treat Kotlin Result<T> as a value-based error container.
- Only mark ThrowsExceptions as "yes" if the function explicitly uses "throw"
  OR calls a function that is clearly not handled by try/catch.

Output format (strict):
ReturnType:
IsSuspend:
ReturnsResult:
ReturnsSuccess:
ReturnsFailure:
ThrowsExceptions:
ThrownTypes:
SideEffects:
ExternalIO:
ErrorHandlingMechanism:
Notes:

Function:
{{FUNCTION}}
`
