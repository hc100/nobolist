<template>
  <div>
    <Nav />
    <header class="bg-white shadow">
      <div class="max-w-7xl mx-auto py-6 px-4 sm:px-6 lg:px-8">
        <h1 class="text-3xl font-bold text-gray-900">パスワードの再設定</h1>
      </div>
    </header>
    <main>
      <div class="max-w-7xl mx-auto py-6 sm:px-6 lg:px-8">
        <div v-if="mode === MODE_INPUT" class="px-4 py-6 sm:px-0">
          <div class="border-2 border-gray-200 rounded h-96">
            <form class="px-8 pt-6 pb-8 mb-4">
              <div
                v-for="(error, index) in errors"
                :key="index"
                class="mb-4 text-red-600"
              >
                {{ error }}
              </div>

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

              <div class="flex items-center justify-between">
                <button
                  class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                  type="button"
                  @click="onSubmit"
                >
                  送信
                </button>
              </div>
            </form>
          </div>
        </div>
        <div v-if="mode === MODE_COMPLETE">
          <div class="border-2 border-gray-200 rounded h-96">
            <div class="m-4">
              ご入力いただいたメールアドレス宛に、お手続きに進めるURLをお送りします
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import ResetPasswordRequest from '~/apollo/resetPasswordRequest'

const MODE_INPUT = 'Input'
const MODE_COMPLETE = 'Complete'

export default {
  apollo: {},
  data() {
    return {
      MODE_INPUT,
      MODE_COMPLETE,
      mode: MODE_INPUT,
      submitting: false,
      errors: [],
      email: null,
    }
  },
  head() {
    return {
      title: 'パスワードの再設定',
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

      if (this.email !== null) {
        this.email = this.email.trim()
      }

      if (!this.email) {
        this.errors.push('メールアドレスを入力してください。')
      }

      if (this.errors.length > 0) {
        return false
      }

      try {
        await this.$apollo
          .mutate({
            mutation: ResetPasswordRequest,
            variables: { email: this.email },
          })
          .then(({ data }) => {
            this.submitting = false
            this.mode = MODE_COMPLETE
            window.scrollTo({
              top: 0,
              behavior: 'smooth',
            })
          })
      } catch (e) {
        console.error(e)
      }
    },
  },
}
</script>

<style></style>
