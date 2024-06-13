export const makeQueryString = (
  params: Record<string, string | number>,
  config: {
    zeroAllow: boolean;
  } = { zeroAllow: false },
) => {
  const searchParams = new URLSearchParams('')
  let validate = (value: string | number): boolean => !!value

  if (config && config.zeroAllow) {
    validate = (value) => value !== undefined && value !== null && value !== ''
  }

  Object.entries(params).forEach(([key, value]) => {
    if (validate(value)) {
      searchParams.append(key, value.toString())
    }
  })

  return searchParams.toString()
}
