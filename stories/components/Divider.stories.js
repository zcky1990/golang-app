import Divider from "../../static/javascript/components/Divider.vue";

export default {
  title: "Components/Divider",
  component: Divider,
  tags: ["autodocs"],
  parameters: {
    docs: {
      description: {
        component: 'A countdown timer component that counts down to a specified date. Customize the title and the end date using props.'
      }
    }
  },
  render: (args) => ({
    components: {
      Divider,
    },
    setup() {
      return {
        ...args,
      };
    },
    template:
      '<Divider :message="message" :message-position="messagePosition" :color="color" :message-transform="messageTransform" />',
  }),
  argTypes: {
    message:{ control: "text" },
    messagePosition:{ 
      options: ['left', 'center', 'right'],
      control: { type: 'select' }, 
    },
    messageTransform:{ 
      options: ['uppercase', 'lowercase', 'capitalize', 'normal-case'],
      control: { type: 'select' }, 
    },
    color:{ 
      options: ['gray', 'black', 'red'],
      control: { type: 'select' }, 
    }
  },
};

export const Center = {
  parameters: {
    docs: {
      description: {
        story: 'Example How Divider component renderer when message position is center'
      }
    }
  },
  args: {
    message:"divider",
    messagePosition:"center",
    messageTransform: "uppercase",
    color:"gray",
  }
}

export const Left = {
  parameters: {
    docs: {
      description: {
        story: 'Example How Divider component renderer when message position is left'
      }
    }
  },
  args: {
    message:"divider",
    messagePosition:"left",
    messageTransform: "uppercase",
    color:"gray",
  },
};


export const Right = {
  parameters: {
    docs: {
      description: {
        story: 'Example How Divider component renderer when message position is right'
      }
    }
  },
  args: {
    message:"divider",
    messagePosition:"right",
    messageTransform: "uppercase",
    color:"gray",
  },
};
