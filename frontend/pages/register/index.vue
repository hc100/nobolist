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
            <div v-if="error" style="color: red">{{ error }}</div>

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
                placeholder="climber@gmail.com"
                required
              />
            </div>
            <div class="flex items-center justify-between">
              <button
                class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
                type="button"
                @click="onSubmit"
              >
                送信する
              </button>
            </div>
          </form>
        </div>
      </div>
      <div v-if="mode === MODE_COMPLETE">
        <div>
          <h2>新規会員登録 メール送信完了</h2>
          <p>
            ご入力いただいたメールアドレス宛に、お手続きに進めるURLをお送りします。
          </p>
        </div>
      </div>
    </main>
  </div>
</template>

<script>
import ExistUserEmail from '~/apollo/existUserEmail'
import CreateUser from '~/apollo/createUser'
import { isValidEmail } from '~/utils/validators.js'

const MODE_INPUT = 'Input'
const MODE_COMPLETE = 'Complete'

export default {
  data() {
    return {
      MODE_INPUT,
      MODE_COMPLETE,
      mode: MODE_INPUT,
      email: null,
      error: null,
      submitting: false,
    }
  },
  methods: {
    async onSubmit(e) {
      e.preventDefault()
      this.error = null
      this.submitting = true

      if (!this.email) {
        this.submitting = false
        this.error = 'メールアドレスを入力してください'
        return
      }

      if (!isValidEmail(this.email)) {
        this.submitting = false
        this.error = 'メールアドレスを正しく入力してください'
        return
      }

      try {
        await this.$apollo
          .query({
            query: ExistUserEmail,
            variables: { email: this.email },
          })
          .then(({ data }) => {
            if (!data.existUserEmail) {
              this.$apollo
                .mutate({
                  mutation: CreateUser,
                  variables: {
                    email: this.email,
                  },
                })
                .then(({ data }) => {
                  this.submitting = false
                  this.mode = MODE_COMPLETE
                  window.scrollTo({
                    top: 0,
                    behavior: 'smooth',
                  })
                })
            } else {
              this.submitting = false
              this.error = 'メールアドレスは既に登録されています'
            }
          })
        this.submitting = false
      } catch (e) {
        this.submitting = false
        this.error = e
      }
    },
  },
}
</script>
