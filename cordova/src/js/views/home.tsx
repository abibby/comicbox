import SeriesList from 'js/components/series-list'
import Layout from 'js/views/layout'
import { Component, h } from 'preact'

export default class Home extends Component {

    public render() {

        return <Layout backLink='/'>
            <h1>Home</h1>
            <SeriesList list='READING' />
        </Layout >
    }
}