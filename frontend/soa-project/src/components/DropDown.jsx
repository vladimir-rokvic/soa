import React from "react"

const DropDownContext = React.createContext({
	open: false,
	setOpen: () => {}
});

const DropDown = ({children, ...props}) => {
  const [open, setOpen] = React.useState(false);
  return (
     <DropdownContext.Provider value={{ open, setOpen }}>
       <div className="relative">{children}</div>
     </DropdownContext.Provider>
  );
};
