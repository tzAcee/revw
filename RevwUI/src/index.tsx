import { render } from 'preact';
import { LocationProvider, Router, Route } from 'preact-iso';
import { GeistProvider, CssBaseline } from '@geist-ui/core'

import { Home } from './pages/Home/index.jsx';
import { NotFound } from './pages/_404.jsx';
import { Header } from './components/header.js';

export function App() {
	return (
		<LocationProvider>
			<GeistProvider themeType="dark">
    			<CssBaseline /> 
				<Header></Header>
				<main>
					<Router>
						<Route path="/" component={Home} />
						<Route default component={NotFound} />
					</Router>
				</main>
			</GeistProvider>
		</LocationProvider>
	);
}

render(<App />, document.getElementById('app'));
