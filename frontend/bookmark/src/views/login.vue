<template>
    <v-app app>
        <v-main class="grey lighten-3">
            <v-alert :type="alertMessage.alertType" class="mx-auto" :style="alertMessage.alertStyle">{{alertMessage.text}}</v-alert>
            <v-card elevation="2" class="mx-auto my-6 pa-5" width=330 height=450>
                <v-form ref="formRef" v-model="valid">
                    <div v-if="status" class="grey--text">注册</div>
                    <div v-else class="grey--text">登录</div>
                    <v-text-field label="username" required v-model="formData.username" :rules="nameRules" :messages="messages.nameErrMessage">
                    </v-text-field>
                    <v-text-field label="E-mail" v-if="status" required v-model="formData.email" :rules="emailRules">
                    </v-text-field>
                    <v-row v-if="status" class="py-0">
                        <v-col class="py-0">
                            <v-text-field label="验证码" required v-model="formData.veriCode">
                            </v-text-field>
                        </v-col>
                        <v-col class="py-0">
                            <v-btn color="primary" small class="ml-2 mt-5" @click="getVerificationCode">发送验证码</v-btn>
                            <p class="mt-1 ml-5 blue--text small">{{messages.codeSendMessage}}</p>
                        </v-col>
                    </v-row>
                    <v-text-field label="password" type="password" required v-model="formData.password" :rules="pwdRules">
                    </v-text-field>
                    <v-text-field label="confirm password" v-if="status" type="password" required v-model="formData.confirmPassword" @blur="checkSame" :rules="pwdRules" :messages="messages.pwdErrMessage">
                    </v-text-field>
                    <v-row style="position:absolute; left:30px; top:360px;">
                        <v-col cols=5>
                            <v-btn color="primary" class="mr-4 mt-5" @click="register" v-if="status">注册</v-btn>
                            <v-btn color="primary" class="mr-4 mt-5" v-else @click="login">登录</v-btn>
                        </v-col>
                        <v-col cols=7 class="pr-0">
                            <div class="mt-6 ml-4 grey--text" style="cursor:pointer;" v-if="status" @click="setStatus(0)">已有账号？登录</div>
                            <div class="mt-6 ml-4 grey--text" style="cursor:pointer;" v-else @click="setStatus(1)">没有账号？注册</div>
                        </v-col>
                    </v-row>
                </v-form>
            </v-card>
        </v-main>
    </v-app>
</template>

<script>
export default {
    data() {
        return {
            alertMessage: {
                alertStyle: {
                    width: "300px",
                    height: "50px",
                    visibility: "hidden"
                },
                alertType: "success",
                text: ""
            },
            messages: {
                pwdErrMessage: "",
                nameErrMessage: "",
                codeSendMessage: ""
            },
            valid: true,
            status: 1, //0为登录，1为注册
            formData: {
                username: "",
                email: "",
                password: "",
                confirmPassword: "",
                veriCode: ""
            },
            nameRules: [v => !!v || "username is required"],
            emailRules: [v => !!v || 'E-mail is required',
                         v => /.+@.+\..+/.test(v) || 'E-mail must be valid'],
            pwdRules: [v => !!v || "password is required",
                       v => v.length > 5 || "password is too short"],
        }
    },

    methods: {
        checkSame() {
            if(this.formData.password != this.formData.confirmPassword) {
                this.messages.pwdErrMessage = "两次输入密码不一致"
                return false
            }
            else {
                this.messages.pwdErrMessage = ""
                return true
            }   
        },

        async getVerificationCode() {
            const {data: data} = await this.$http.post('verificationCode', {
                username: this.formData.username,
                email: this.formData.email
            })
            if(data.code != 200) {
                this.messages.nameErrMessage = data.msg
            }
            else {
                this.messages.nameErrMessage = ""
                this.messages.codeSendMessage = "验证码已发送"
            }
        },
        async register() {
            if(this.$refs.formRef.validate() && this.checkSame()) {
                const {data: data} = await this.$http.post(`register/${this.formData.veriCode}`, {
                    username: this.formData.username,
                    password: this.formData.password,
                    email: this.formData.email
                })
                //注册成功
                if(data.code == 200) {
                    this.alertMessage.alertStyle.visibility = "visible"
                    this.alertMessage.alertType = "success",
                    this.alertMessage.text = "注册成功"
                }
                else {
                    this.alertMessage.alertStyle.visibility = "visible"
                    this.alertMessage.alertType = "error",
                    this.alertMessage.text = data.msg
                }
            }
        },

        async login() {
            const {data: data} = await this.$http.post('login', {
                username: this.formData.username,
                password: this.formData.password
            })
            if(data.code == 200) {
                    this.alertMessage.alertStyle.visibility = "visible"
                    this.alertMessage.alertType = "success",
                    this.alertMessage.text = "登录成功"
                    window.localStorage.setItem("token", data.token)
                    this.$router.push('/space')
                }
                else {
                    this.alertMessage.alertStyle.visibility = "visible"
                    this.alertMessage.alertType = "error",
                    this.alertMessage.text = data.msg
                }
        },

        setStatus(sta) {
            this.status = sta
        }
    }
}
</script>

<style scoped>
    .small {
        font-size: 10px;
    }
</style>