# Antrea LFX Mentorship Test Task – Renovate Playground  
> **Project ID:** Replace Dependabot with Renovate for automatic dependency updates (#7155)

---

## 1 ▪ Overview
This repository is a **self-contained Go playground** that shows how `Renovate` can:

1. Detect an **out-of-date dependency** with a **known Critical/High CVE**  
2. Open an **automated pull request** pinning the dependency to the first patched version  
3. Provide a **Dependency Dashboard** so maintainers can track outstanding security upgrades

The repo was specifically built to **demonstrate Renovate on a Go module**, as requested by the test-task instructions.

---

## 2 ▪ Vulnerable Dependency Chosen

| Item | Details |
|------|---------|
| **Library** | `github.com/dgrijalva/jwt-go` |
| **Pinned Version** | `v3.2.0` |
| **Vulnerability** | **CVE-2020-26160** – JWT tokens signed with RSA / ECDSA can be spoofed ([NVD link](https://nvd.nist.gov/vuln/detail/CVE-2020-26160)) |
| **Fixed Version** | `v3.2.1` (or migration to `github.com/golang-jwt/jwt`) |

> **Why this lib?**  
> *jwt-go* has been used by many CNCF projects in the past (including older versions of Antrea), so it’s a realistic scenario for supply-chain risk.

---

## 3 ▪ Program Description (`main.go`)
The Go program is intentionally minimal:

```go
package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go" // vulnerable version v3.2.0
)

func main() {
	token := jwt.New(jwt.SigningMethodHS256)
	fmt.Println("Generated JWT:", token)
}
Running the app is not intended to exploit the CVE; it simply imports and uses the vulnerable package so Renovate flags it.

4 ▪ Renovate Configuration (renovate.json)
json
Copy
Edit
{
  "extends": ["config:base"],
  "dependencyDashboard": true,
  "packageRules": [
    {
      "matchManagers": ["gomod"],
      "vulnerabilityAlertsOnly": true          // ONLY open PRs if CVE present
    }
  ],
  "automerge": true,
  "labels": ["dependencies", "security"]
}
Key points
vulnerabilityAlertsOnly:true  → the task requirement to update only when a vulnerability exists

dependencyDashboard:true  → single-pane view for maintainers

automerge:true  → secure point-fixes can land automatically after CI

5 ▪ How to Reproduce Locally 💻
bash
Copy
Edit
# clone & run
git clone https://github.com/ChaithraDevops/antrea-renovate-lfx.git
cd antrea-renovate-lfx
go run main.go
You should see a Generated JWT: message – proof the vulnerable lib is loaded.

6 ▪ How Renovate Demonstrates Its Job 🤖
Renovate GitHub App is installed on this repo.

On first scan it identifies jwt-go v3.2.0 → CVE-2020-26160.

It opens a PR titled “chore(deps): update module dgrijalva/jwt-go to v3.2.1”.

The PR carries the GitHub Security Advisory details + a changelog snippet.

If CI passes, automerge merges and closes the vulnerability.

(See the Pull Requests tab for live evidence.)

7 ▪ DevOps / Supply-Chain-Security Value
Pain-point	Renovate Solution
Manual CVE chasing	Automated PRs for vulnerable versions
Forgotten release branches	Config can target release-* branches too
Busy review queues	Dependency Dashboard + labels + automerge

This directly supports the Antrea objective of stronger and more automated dependency hygiene.

8 ▪ Author & Attribution
Chaithra L U (GitHub: ChaithraDevops, LF Username: Chaithra Mallik)
Submission for LFX Mentorship 2025 Term 2 – Antrea Project.
