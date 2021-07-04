type LayoutProps = {
  children: JSX.Element
}

const Layout: React.FC<LayoutProps> = ({ children }): JSX.Element => (
  <div>
    {children}
  </div>
)


export default Layout
