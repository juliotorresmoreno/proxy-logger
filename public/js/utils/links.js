
function attachLinks(component) {
    const minPath = (location.protocol + '//' + location.host).length;
    const links = component.querySelectorAll('a');
    links.forEach((link) => {
        link.onclick = (evt) => {
            evt.preventDefault();
            links.forEach((currentLink) => {
                let className = currentLink.className;
                className = className.replace(/\sactive/, '');
                currentLink.className = className;
            });
            link.className += ' active';
            history.pushState(undefined, undefined, link.href);
            appStore.setState({
                route: link.href.substr(minPath)
            });
        }
    });
}