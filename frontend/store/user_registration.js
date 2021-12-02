export const state = () => ({
  email: null,
  key: null,
})

export const mutations = {
  setEmail(state, email) {
    state.email = email
  },
  setKey(state, key) {
    state.key = key
  },
}
