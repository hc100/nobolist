<template>
  <div>
    <Header />
    <main class="pl-8">
      <div class="w-full max-w-xs">
        <div class="bg-green-200 rounded px-6 pt-6 pb-6 mb-4">ログイン</div>
      </div>
      <div class="w-full max-w-xs">
        <form class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">
          <p v-for="(error, index) in errors" :key="index" class="required">
            {{ error }}
          </p>

          <div class="mb-4">
            <label
              class="block text-gray-700 text-sm font-bold mb-2"
              for="email"
            >
              メールアドレス
            </label>
            <input
              id="email"
              v-model="email"
              class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
              type="text"
              placeholder="climber@nobolist.jp"
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
              required
            />
          </div>
          <div class="flex items-center justify-between">
            <button
              class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
              type="button"
              @click="onSubmit"
            >
              ログイン
            </button>
          </div>

          <div class="mt-4">
            <a class="underline" href="/reset_password/"
              >パスワードを忘れた方はこちら</a
            >
          </div>

          <div class="mt-4">
            <a class="underline" href="/register/">新規会員登録はこちら</a>
          </div>
        </form>
      </div>
    </main>
  </div>
</template>

<script>
import Login from '~/apollo/login'

export default {
  apollo: {},
  data() {
    return {
      submitting: false,
      errors: [],
      email: null,
      password: null,
    }
  },
  head() {
    return {
      title: 'ログイン',
    }
  },
  computed: {},
  watch: {},
  beforeMount() {},
  mounted() {},
  created() {},
  methods: {
    async onSubmit(e) {
      e.preventDefault()
      this.errors = []
      this.submitting = true

      if (this.password !== null) {
        this.password = this.password.trim()
      }
      if (this.email !== null) {
        this.email = this.email.trim()
      }

      if (!this.email) {
        this.errors.push('メールアドレスを入力してください。')
      }

      if (!this.password) {
        this.errors.push('パスワードを入力してください。')
      }

      if (this.errors.length > 0) {
        return false
      }

      const credentials = {
        email: this.email,
        password: this.password,
      }

      try {
        const res = await this.$apollo
          .mutate({
            mutation: Login,
            variables: credentials,
          })
          .then(({ data }) => data && data.login)
        await this.$apolloHelpers.onLogin(res.accessToken)
        this.$router.push({ name: 'dashboard' })
      } catch (e) {
        console.error(e)
      }
    },
  },
}
</script>

<style></style>
