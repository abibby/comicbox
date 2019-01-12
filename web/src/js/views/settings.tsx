import autobind from 'autobind-decorator'
import * as s from 'css/settings.scss'
import { logout, user } from 'js/auth'
import { gql } from 'js/graphql'
import User from 'js/model/user'
import url from 'js/url'
import Layout from 'js/views/layout'
import { Component, h } from 'preact'
import Button from 'preact-material-components/Button'
import TextField from 'preact-material-components/TextField'
import { Link } from 'preact-router'
import { OpenForm } from 'js/components/modal';

/**
 * Things settings will do.
 *
 * - Change current user's name, username and password
 * - logout
 * - start scan
 * - admin stuff
 *   - add users
 *   - delete users
 *   - change user roles
 * - change things from the config file?
 *   - book folder
 */

interface State {
    address: string

    me: User
    username: string
    name: string

    oldPass: string
    newPass: string
    repeatNewPass: string

    // Admin related vars
    newNameToAdd: string
    newUsernameToAdd: string
    newUserPasswordToAdd: string
    userToDelete: string
    userToMakeAdmin: string
    userToRevokeAdmin: string
}

export default class Settings extends Component<{}, State> {

    public componentDidMount() {
        user().then(me => {
            this.setState({
                me: me,
                username: me.username,
                name: me.name,
            })
        })
    }

    public render() {

        return <Layout backLink='/'>
            <h1>Settings</h1>
            <Button onClick={this.btnTest}>Test</Button>
        </Layout >
    }

    @autobind
    private async btnTest() {
        const data = OpenForm({ title: 'Username' }, <div>
            <TextField label='Username' />
        </div>)
        console.log(data)
    }

    @autobind
    private async btnUpdateUser() {
        const me = this.state.me

        me.name = this.state.name
        me.username = this.state.username

        await me.save()

        this.setState({
            me: me,
            username: me.username,
            name: me.name,
        })
    }

    @autobind
    private async btnUpdatePassword() {
        const me = this.state.me

        // TODO: Add check to match original old password with given old password for security
        if (this.state.newPass !== this.state.repeatNewPass) {
            alert("Your passwords don't match. Please try again.")
            return
        }

        me.password = this.state.newPass

        await me.save()

        this.setState({
            newPass: '',
            repeatNewPass: '',
        })
    }

    @autobind
    private btnLogout() {
        logout()
    }

    @autobind
    private async btnScan() {
        const response = await fetch(await url('/api/scan'))
        if (!response.ok) {
            // TODO: something here
        }
    }

    // Adds a user to the database
    @autobind
    private async btnAddUser() {
        // Returns whats in the top block (id,name,username)
        const response = await gql(`
            new_user(data: $user) {
                id
                name
                username
            }`, {
                user: 'UserInput!',
            }, {
                user: {
                    name: this.state.newNameToAdd,
                    username: this.state.newUsernameToAdd,
                    password: this.state.newUserPasswordToAdd,
                },
            }, true)
    }

    @autobind
    private async btnDeleteUser() {
        // TODO: Add functionality to delete user from the database
    }

    @autobind
    private async btnGrantAdminStatus() {
        // TODO: Add granting admin status functionality once user groups are implemented
    }

    @autobind
    private async btnRevokeAdminStatus() {
        // TODO: Add revoking admin status functionality once user groups are implemented
    }
}
