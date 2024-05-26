package policy

#Policy: {
  Version!: string
  Statement!: [#Statement]
}

#Statement: close({
  Effect!: "Allow" | "Deny"
  Action!: string | [string]
  Resource!: string | [string]
})
