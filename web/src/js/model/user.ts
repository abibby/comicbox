import { Model, prop, table } from 'js/model/model'

@table('users', 'user', 'UserInput!')
export default class User extends Model {

    public static searchTypes = {
        me: { type: 'Boolean' },
    }

    public static async me(): Promise<User> {
        return await User.where('me', true).first()
    }

    @prop('String')
    public id: string

    @prop('String')
    public name: string

    @prop('String')
    public username: string

    @prop('String', { writeOnly: true })
    public password: string

    public get link() {
        return `/user/${this.username}`
    }

    public get sortIndex() {
        return `user-${this.username}`
    }
}