import ButtonAppBar from "../components/AppBar";

export default function Layout({ children }) {
  return (
    <div>
      <ButtonAppBar />
      <main>{children}</main>
    </div>
  );
}
