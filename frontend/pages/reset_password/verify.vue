<template>
  <div>
    <Header />
    <main class="pl-8">
      <div v-if="mode === MODE_INPUT">
        <div class="w-full max-w-xs">
          <div class="bg-green-200 rounded px-6 pt-6 pb-6 mb-4">
            パスワード再設定 入力
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
                for="password"
              >
                新しいパスワード
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
                設定する
              </button>
            </div>
          </form>
        </div>
      </div>

      <div v-if="mode === MODE_COMPLETE">
        <div class="w-full max-w-xs">
          <div class="bg-green-200 rounded px-6 pt-6 pb-6 mb-4">
            パスワード再設定 完了
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import IsValidResetPasswordKey from '~/apollo/isValidResetPasswordKey'
import ResetPassword from '~/apollo/resetPassword'
import { isValidPassword } from '~/utils/validators.js'

const MODE_INPUT = 'Input'
const MODE_COMPLETE = 'Complete'

export default {
  apollo: {},
  async asyncData({ error, route, app, store }) {
    try {
      await app.apolloProvider.defaultClient
        .query({
          query: IsValidResetPasswordKey,
          variables: { key: route.query.key },
        })
        .then(({ data }) => {
          store.commit('reset_password/setKey', route.query.key)
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
      title: 'パスワード再設定',
    }
  },
  computed: {},
  watch: {},
  beforeMount() {},
  mounted() {
    this.mode = MODE_INPUT
    this.key = this.$store.state.reset_password.key
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
          password: this.password,
        }
        await this.$apollo
          .mutate({
            mutation: ResetPassword,
            variables: input,
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
