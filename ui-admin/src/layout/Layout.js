import Navbar from './Navbar';

function Layout({ children }) {
  return (
    <div className="row d-flex flex-wrap">
      <Navbar />
      <main>
        <div className="px-2">
          {children}
        </div>
      </main>
    </div>
  );
}

export default Layout;
