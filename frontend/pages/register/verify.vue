<template>
  <div>
    <Header />
    <main>
      <div v-if="mode === MODE_INPUT">
        <div class="w-full max-w-xs">
          <div class="bg-green-200 rounded px-6 pt-6 pb-6 mb-4">
            新規会員登録 入力
          </div>
        </div>
        <div class="w-full max-w-xs">
          <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
            <p v-for="(error, index) in errors" :key="index" class="required">
              {{ error }}
            </p>

            <div class="mb-4">
              <label
                class="block text-gray-700 text-sm font-bold mb-2"
                for="name"
              >
                名前（表示用）
              </label>
              <input
                id="name"
                v-model="name"
                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                type="text"
                placeholder="名無しのクライマー"
                required
              />
            </div>

            <div class="mb-4">
              <label
                class="block text-gray-700 text-sm font-bold mb-2"
                for="password"
              >
                パスワード
              </label>
              <input
                id="password"
                v-model="password"
                class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
                type="password"
                placeholder="半角英数字の組み合わせ"
                required
              />
            </div>
            <div class="flex items-center justify-between">
              <button
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                type="button"
                @click="onSubmit"
              >
                登録する
              </button>
            </div>
          </form>
        </div>
      </div>

      <div v-if="mode === MODE_COMPLETE">
        <div class="w-full max-w-xs">
          <div class="bg-green-200 rounded px-6 pt-6 pb-6 mb-4">
            新規会員登録 完了
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import IsValidRegistrationKey from '~/apollo/isValidRegistrationKey'
import RegisterUser from '~/apollo/registerUser'
import { isValidPassword } from '~/utils/validators.js'

const MODE_INPUT = 'Input'
const MODE_COMPLETE = 'Complete'

export default {
  apollo: {},
  async asyncData({ error, route, app, store }) {
    try {
      await app.apolloProvider.defaultClient
        .query({
          query: IsValidRegistrationKey,
          variables: { key: route.query.key },
        })
        .then(({ data }) => {
          store.commit(
            'user_registration/setEmail',
            data.isValidRegistrationKey.email
          )
          store.commit('user_registration/setKey', route.query.key)
        })
    } catch (e) {
      console.log(e)
      error({
        statusCode: 404,
      })
    }
  },
  data() {
    return {
      MODE_INPUT,
      MODE_COMPLETE,
      submitting: false,
      errors: [],
      mode: MODE_INPUT,
      email: null,
      key: null,
      password: null,
      name: null,
    }
  },
  head() {
    return {
      title: '会員登録',
    }
  },
  computed: {},
  watch: {},
  beforeMount() {},
  mounted() {
    this.mode = MODE_INPUT
    this.email = this.$store.state.user_registration.email
    this.key = this.$store.state.user_registration.key
  },
  created() {},
  methods: {
    async onSubmit(e) {
      e.preventDefault()
      this.errors = []
      this.submitting = true

      if (this.password !== null) {
        this.password = this.password.trim()
      }
      if (this.name !== null) {
        this.name = this.name.trim()
      }

      if (!this.name) {
        this.errors.push('名前を入力してください。')
      }

      if (!this.password) {
        this.errors.push('パスワードを入力してください。')
      } else if (!isValidPassword(this.password)) {
        this.errors.push(
          'パスワードを半角英数字を含む8-12桁で入力してください。'
        )
      }

      if (this.errors.length > 0) {
        return false
      }

      try {
        const input = {
          key: this.key,
          name: this.name,
          password: this.password,
        }
        await this.$apollo
          .mutate({
            mutation: RegisterUser,
            variables: { input },
          })
          .then(({ data }) => {
            this.submitting = false

            window.scrollTo({
              top: 0,
              behavior: 'smooth',
            })

            this.mode = MODE_COMPLETE
          })
      } catch (e) {
        console.log(e)
        this.submitting = false
        this.errors.push(e)

        window.scrollTo({
          top: 0,
          behavior: 'smooth',
        })

        this.mode = MODE_INPUT
      }
    },
  },
}
</script>

<style></style>
