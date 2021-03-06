import { HamburgerIcon } from "@chakra-ui/icons"
import {
  Avatar,
  Box,
  Flex,
  FlexProps,
  HStack,
  IconButton,
  Menu,
  MenuButton,
  MenuDivider,
  MenuItem,
  MenuList,
  Text,
  useColorModeValue,
  useToast
} from "@chakra-ui/react"
import { useCookies } from "react-cookie"
import { GrGamepad, GrList } from "react-icons/gr"
import { Link } from "react-router-dom"
import { makeToastData } from "../../helpers"


type LayoutProps = {
  children: JSX.Element
}


const MobileNav = ({ ...rest }: FlexProps) => {
  const [cookies, , removeCookie] = useCookies(["access", "refresh"])
  const toast = useToast()

  const isAuthenticated: boolean = !!cookies.access && !!cookies.refresh

  return (
    <Flex
      ml={{ base: 0, md: 0 }}
      px={{ base: 4, md: 4 }}
      height="20"
      alignItems="center"
      bg={useColorModeValue("white", "gray.900")}
      borderBottomWidth="1px"
      borderBottomColor={useColorModeValue("gray.200", "gray.700")}
      justifyContent={{ base: "space-between" }}
      {...rest}>

      <Menu>
        <MenuButton
          as={IconButton}
          aria-label='Options'
          icon={<HamburgerIcon />}
          variant='outline'
        />
        <MenuList>
          <Link to="/">
            <MenuItem icon={<GrGamepad />}>
              Game
            </MenuItem>
          </Link>
          <Link to="/games">
            <MenuItem icon={<GrList />}>
              Game List
            </MenuItem>
          </Link>
        </MenuList>
      </Menu>

      <Text
        display={{ base: "flex" }}
        fontSize="2xl"
        fontFamily="monospace"
        fontWeight="bold">
        <Link to="/">
          animated
        </Link>
      </Text>

      <HStack spacing={{ base: "0", md: "6" }}>
        {/* <IconButton
          size="lg"
          variant="ghost"
          aria-label="open menu"
          icon={<FiBell />}
        /> */}
        <Flex alignItems={"center"}>
          <Menu>
            <MenuButton
              py={2}
              transition="all 0.3s"
              _focus={{ boxShadow: "none" }}>

              <HStack>
                <Avatar
                  size={"sm"}
                  src={`${window.location.origin}/img/default-avatar.png`}
                />
                {/* <VStack
                  display={{ base: "none", md: "flex" }}
                  alignItems="flex-start"
                  spacing="1px"
                  ml="2">
                  <Text fontSize="sm">Justina Clark</Text>
                  <Text fontSize="xs" color="gray.600">
                    User
                  </Text>
                </VStack>
                <Box display={{ base: "none", md: "flex" }}>
                  <FiChevronDown />
                </Box> */}
              </HStack>

            </MenuButton>
            <MenuList
              bg={useColorModeValue("white", "gray.900")}
              borderColor={useColorModeValue("gray.200", "gray.700")}>
              <Link to="/login">
                <MenuItem>Profile</MenuItem>
              </Link>
              {isAuthenticated && <>
                <MenuDivider />
                <MenuItem onClick={() => {
                  removeCookie("access")
                  removeCookie("refresh")
                  toast(makeToastData({ title: "Logged out" }))
                }}>Sign out</MenuItem>
              </>}
            </MenuList>
          </Menu>
        </Flex>
      </HStack>
    </Flex>
  )
}

const Layout: React.FC<LayoutProps> = ({ children }): JSX.Element => {
  // const { isOpen, onOpen, onClose } = useDisclosure()
  return (
    <Box minH="100vh">
      <MobileNav />
      <Box ml={{ base: 0, md: 0 }} p="4">
        {children}
      </Box>
    </Box>
  )

}

export default Layout
