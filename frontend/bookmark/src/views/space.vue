<template>
    <v-app>
        <v-main class="brown lighten-5">
            <div class="grey--text ml-12 mt-5 mb-0" style="font-size: 20px; display: inline-block;">分类</div>

            <v-dialog v-model="categoryDialog" persistent max-width="600px">
                <template v-slot:activator="{ on, attrs }">
                    <v-btn color="primary" fab dark v-bind="attrs" v-on="on" class="pa-0 ml-6">
                        <v-icon size="40" class="ma-0" style="cursor: pointer;">{{'mdi-plus'}}</v-icon>
                    </v-btn>
                </template>
                <v-card class="py-2 px-5">
                    <v-card-title>
                        <span class="text-h5">新建标签</span>
                    </v-card-title> 
                    <v-text-field label="标签名称" required v-model="newCategory"></v-text-field>
                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" text @click="categoryDialog = false">关闭</v-btn>
                        <v-btn color="blue darken-1" text @click="addCategory()">提交</v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
            
            <div v-if="currentCategory != 0" style="float: right; font-size: 20px; cursor: pointer;" 
                class="mr-12 mt-5 grey--text" 
                @click="backToParent()">回到上级</div>
            <v-sheet class="sheet brown lighten-5">
                <v-card v-for="item in categories" :key="item.ID" 
                width=200 height=240 class="category" @click="changeCategory(item.ID)">
                    <v-img width=200 height=130 :src=item.pic></v-img>
                    <v-card-text class="pa-0">
                        <v-row class="ma-0 pa-0">
                            <v-col cols=8><div>{{item.name}}</div></v-col>
                            <v-col cols=2 class="px-0">
                                <v-dialog v-model="bookmarkDialog2" persistent max-width="600px">
                                    <template v-slot:activator="{ on, attrs }">
                                        <v-icon  class="ma-0" v-bind="attrs" v-on="on" @click="selectCid(item.ID)">{{'mdi-plus'}}</v-icon>
                                    </template>
                                    <v-card>
                                        <v-card-title>
                                            <span class="text-h5">新建书签</span>
                                        </v-card-title> 
                                        <v-text-field label="书签URL" required v-model="newBookmarkUrl"></v-text-field>  
                                        <v-card-actions>
                                            <v-spacer></v-spacer>
                                            <v-btn color="blue darken-1" text @click="bookmarkDialog2 = false">关闭</v-btn>
                                            <v-btn color="blue darken-1" text @click="addBookmark()">提交</v-btn>
                                        </v-card-actions>
                                    </v-card>
                                </v-dialog>
                            </v-col>
                            <v-col cols=2 class="pa-0">
                                <v-menu offset-y>
                                    <template v-slot:activator="{ on, attrs }">
                                         <v-icon class="pa-0 mt-3" v-bind="attrs" v-on="on">{{'mdi-trash-can-outline'}}</v-icon>
                                    </template>
                                    <v-list class="py-0">
                                        <v-list-item class="pa-1 green--text" style="cursor: pointer;" @click="deleteCategory(item.ID, true)">仅删除标签</v-list-item>
                                        <v-list-item class="pa-1 red--text" style="cursor: pointer;" @click="deleteCategory(item.ID, false)">删除标签及内容</v-list-item>
                                    </v-list>
                                 </v-menu>
                            </v-col>
                        </v-row>
                        <p class="grey--text small mt-1 mb-0 ml-3">创建时间</p>
                        <p class="grey--text small mt-1 ml-3">书签总数</p>
                    </v-card-text>
                </v-card>
            </v-sheet>
            <v-divider class="mt-5 mb-0"></v-divider>
            <div class="grey--text ml-12 mt-5 mb-0" style="font-size: 20px; display: inline-block;">书签</div>

            <v-dialog v-model="bookmarkDialog" persistent max-width="600px">
                <template v-slot:activator="{ on, attrs }">
                    <v-btn color="primary" fab dark v-bind="attrs" v-on="on" class="pa-0 ml-6">
                        <v-icon size="40" class="ma-0" style="cursor: pointer;">{{'mdi-plus'}}</v-icon>
                    </v-btn>
                </template>
                <v-card>
                    <v-card-title>
                        <span class="text-h5">新建书签</span>
                    </v-card-title> 
                    <v-text-field label="书签URL" required v-model="newBookmarkUrl"></v-text-field>  
                    <v-card-actions>
                        <v-spacer></v-spacer>
                        <v-btn color="blue darken-1" text @click="bookmarkDialog = false">关闭</v-btn>
                        <v-btn color="blue darken-1" text @click="addBookmark()">提交</v-btn>
                    </v-card-actions>
                </v-card>
            </v-dialog>
            <v-btn style="margin-left:786px;" @click="batchOperation()">批量操作</v-btn>
            <v-btn class="ml-6" @click="cut()" :disabled="!enableCut">剪切</v-btn>
            <v-btn class="ml-6" @click="paste()" :disabled="!enablePaste">粘贴</v-btn>
            <v-sheet class="brown lighten-5 px-5">
                <v-card width="1000" height="55" class="mt-5 ml-7"
                        v-for="(item, index) in bookmarks" :key="item.ID" style="cursor: pointer;">
                    <v-row class="pa-0 ma-0 grey--text">
                        <v-col class="pa-0 ma-0" cols=1 @click="navigate(item.url)">
                            <img width=55 height=55 :src="'data:image/png;base64,'+item.icon"/>
                        </v-col>
                        <v-col class="pa-0 ma-0 hidden" cols=5 align-self="center" @click="navigate(item.url)">
                            <div>{{item.title}}</div>
                        </v-col>
                        <v-col class="pa-0 ma-0 hidden" cols=4 align-self="center" @click="navigate(item.url)">
                            <div>{{item.url}}</div>
                        </v-col>
                        <v-col cols=1 class="pl-8 pt-4">
                            <v-icon @click="deleteBookmark(item.ID)">{{'mdi-trash-can-outline'}}</v-icon>
                        </v-col>
                        <v-col cols=1 class="pa-0 ma-0">
                            <v-checkbox class="mt-4 ml-6 pa-0" v-model="currentBatchSelect[index]" :disabled="!enableBatchSelect" @click="select(item.ID)"></v-checkbox>
                        </v-col>
                    </v-row>
                </v-card>
            </v-sheet>
        </v-main>
    </v-app>
</template>

<script>
export default {
    data() {
        return {
            bookmarkDialog2: false,
            bookmarkDialog: false,
            categoryDialog: false,

            parentCategory: [],
            currentCategory: 0,

            selected: new Map(),
            clipBoard: [],
            batchSelects: new Map(),
            currentBatchSelect: [],

            categories: [],
            bookmarks: [],

            newCategory: "",
            newBookmarkUrl: "",

            selectedCid: 0,

            enableBatchSelect: false,
            enableCut: false,
            enablePaste: false,

            selectedNum: 0
        }
    },
    created() {
        this.getCategories()
        this.getBookmarks()
    },
    methods: {
        select(id) {
            if(!this.selected.has(id)) {
                this.selected.set(id, 0)
                this.selectedNum++
            }  
            else {
                this.selected.delete(id)
                this.selectedNum--
            }
            if(this.selectedNum > 0)
                this.enableCut = true
            else 
                this.enableCut = false
        },
        cut() {
            this.selected.forEach((value, key) => {
                this.clipBoard.push(key)
            })
            this.enablePaste = true
            this.enableCut = false
        },
        async paste() {
            const {data: data} = await this.$http.put('bookmark/', {
                ids: this.clipBoard,
                cid: this.currentCategory
            })
            this.selected = new Map()
            this.batchSelects = new Map()
            this.currentBatchSelect = []
            this.clipBoard = []
            this.getBookmarks()

            this.enableBatchSelect = false
            this.enablePaste = false
        },
        batchOperation() {
            this.enableBatchSelect = !this.enableBatchSelect
            this.selected = new Map()
            this.batchSelects = new Map()
            this.currentBatchSelect = []
            this.clipBoard = []
            this.currentBatchSelect = []
            for(let i = 0; i < this.bookmarks.length; i++)
                this.currentBatchSelect.push(false)
        },

        async getCategories() {
            const {data: data} = await this.$http.get(`categories/${this.currentCategory}`)
            if(data.code == 200) {
                this.categories = data.data
            }
        },
        async getBookmarks() {
            this.done = false
            const {data: data} = await this.$http.get(`bookmark/${this.currentCategory}`)
            if(data.code == 200) {
                this.bookmarks = data.data
                this.currentBatchSelect = []
                for(let i = 0; i < this.bookmarks.length; i++)
                     this.currentBatchSelect.push(false)
            }
        },
        navigate(url) {
            window.open(url)
        },
        async changeCategory(cid) {
            this.parentCategory.push(this.currentCategory)
            let tmp = this.currentCategory
            let tmpSelect = this.currentBatchSelect
            this.currentCategory = cid
            this.getCategories()
            await this.getBookmarks()

            this.batchSelects.set(tmp, tmpSelect)
            if(this.batchSelects.get(cid) != undefined) {
                this.currentBatchSelect = this.batchSelects.get(cid)
            }
        },
        async backToParent() {
            let cid = this.currentCategory
            let tmpSelect = this.currentBatchSelect
            this.currentCategory = this.parentCategory.pop()
            this.getCategories()
            await this.getBookmarks()

            this.batchSelects.set(cid, tmpSelect)
            if(this.batchSelects.get(this.currentCategory) != undefined) {
                this.currentBatchSelect = this.batchSelects.get(this.currentCategory)
            }
        },
        async addCategory() {
            await this.$http.post('category', {
                pid: this.currentCategory,
                name: this.newCategory
            })
            this.getCategories()
            this.categoryDialog = false
            this.newCategory = ""
        },
        async deleteCategory(id, RetainContent) {
            await this.$http.delete(`category/${id}`, {
                data: {
                    ids: [],
                    retain_content: RetainContent
                }   
            })
            this.getCategories()
            this.getBookmarks()
        },
        async deleteBookmark(id) {
            await this.$http.delete('bookmark', {
                data: {
                    ids: [id]
                }
            })
            this.getBookmarks()
        },
        async addBookmark() {
            let c = this.currentCategory
            if(this.selectedCid != 0)
                c = this.selectedCid
            await this.$http.post('bookmark', {
                url: this.newBookmarkUrl,
                cid: c
            })
            this.newBookmarkUrl = ""
            this.bookmarkDialog = false
            this.bookmarkDialog2 = false
            this.selectedCid = 0
            this.getBookmarks()
        },
        selectCid(id) {
            this.selectedCid = id
        }
    }
}
</script>

<style scoped>
    .sheet {
        padding: 10px;
        display: flex;
        flex-wrap: wrap;
    }
    .category {
        margin-top: 20px;
        margin-left: 40px;
    }
    .small {
        font-size: 10px;
    }
    .hidden {
        font-style: italic;
        -webkit-line-clamp: 2; /*（用来限制在一个块元素显示的文本的行数，2 表示最多显示 2 行。为了实现该效果，它需要组合其他的 WebKit 属性）*/
        display: -webkit-box; /*（和 1 结合使用，将对象作为弹性伸缩盒子模型显示 ）*/
        -webkit-box-orient: vertical; /*（和 1 结合使用 ，设置或检索伸缩盒对象的子元素的排列方式 ） */
        overflow: hidden; /*（文本溢出限定的宽度就隐藏内容） */
        text-overflow: ellipsis; /*（多行文本的情况下，用省略号 “…” 隐藏溢出范围的文本)*/
    }
</style>